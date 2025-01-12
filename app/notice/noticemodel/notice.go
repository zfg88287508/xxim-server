package noticemodel

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xredis"
	"github.com/cherish-chat/xxim-server/common/xredis/rediskey"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type (
	Notice struct {
		// 通知id 由convId+会话自增id+userId组成 主键
		NoticeId string `gorm:"column:noticeId;type:char(152);primary_key;not null;" json:"noticeId"`
		// 会话id
		ConvId string `gorm:"column:convId;type:char(96);not null;index:cu;index:unique_in_conv,unique;" json:"convId"`
		// 自增id
		ConvAutoId int64 `gorm:"column:convAutoId;type:bigint(20);not null;index;" json:"convAutoId"`
		// 接收人id 如果为空表示广播
		UserId   string `gorm:"column:userId;type:char(32);not null;index:cu;" json:"userId"`
		UniqueId string `gorm:"column:uniqueId;type:char(160);not null;index:unique_in_conv,unique;" json:"uniqueId"`

		Options NoticeOption `gorm:"column:options;type:json;" json:"options"`
		// 创建时间
		CreateTime int64 `gorm:"column:createTime;type:bigint(13);not null;default:0;" json:"createTime"`
		// ContentType and Content
		ContentType int32  `gorm:"column:contentType;type:int(11);not null;default:0;" json:"contentType"`
		Content     []byte `gorm:"column:content;type:blob;not null;" json:"content"`
		// 标题
		Title string `gorm:"column:title;type:varchar(255);not null;default:'';" json:"title"`
		// 扩展字段
		Ext []byte `gorm:"column:ext;type:blob;not null;" json:"ext"`
		// 查询用户没有消费的通知时，先获取会话最大ConvAutoId，使用索引查询 where ((convId=? and userId=?) or (convId=? and userId="")) and convAutoId > ? order by convAutoId asc limit 1000
	}
	NoticeMaxConvAutoId struct {
		ConvId     string `gorm:"column:convId;type:char(96);not null;primary_key;" json:"convId"`
		ConvAutoId int64  `gorm:"column:convAutoId;type:bigint(20);not null;" json:"convAutoId"`
	}
	NoticeOption struct {
		StorageForClient bool `gorm:"column:storageForClient;type:tinyint(1);not null" json:"storageForClient"`
		UpdateConvNotice bool `gorm:"column:updateConvNotice;type:tinyint(1);not null" json:"updateConvNotice"`
	}
	NoticeAckRecord struct {
		ConvId     string `gorm:"column:convId;type:char(96);index:cud,unique;default:'';" json:"convId"`
		UserId     string `gorm:"column:userId;type:char(32);index:cud,unique;default:'';" json:"userId"`
		DeviceId   string `gorm:"column:deviceId;type:char(96);index:cud,unique;default:'';" json:"deviceId"`
		ConvAutoId int64  `gorm:"column:convAutoId;type:bigint(20);not null;" json:"convAutoId"`
	}
)

func (m *Notice) TableName() string {
	return "notice"
}

func (m *NoticeMaxConvAutoId) TableName() string {
	return "notice_max_conv_auto_id"
}

func (m *NoticeAckRecord) TableName() string {
	return "notice_ack_record"
}

func (m NoticeOption) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *NoticeOption) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), m)
}

func (m *Notice) Insert(ctx context.Context, tx *gorm.DB, rc *redis.Redis) error {
	if m.ConvAutoId == 0 {
		// 获取maxConvAutoId
		convAutoId, err := GetMaxConvAutoId(ctx, rc, m.ConvId, 1)
		if err != nil {
			return err
		}
		m.ConvAutoId = convAutoId
	}
	if m.NoticeId == "" {
		m.NoticeId = pb.ServerNoticeId(m.ConvId, m.ConvAutoId, m.UserId)
	}
	if m.CreateTime == 0 {
		m.CreateTime = time.Now().UnixMilli()
	}
	if m.Ext == nil {
		m.Ext = make([]byte, 0)
	}
	if m.Content == nil {
		m.Content = make([]byte, 0)
	}
	// upsert insert on duplicate key
	return tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "uniqueId"}, {Name: "convId"}},
		// map 更新
		DoUpdates: clause.Assignments(map[string]interface{}{
			// 更新 noticeId
			"noticeId": m.NoticeId,
			// 更新 convAutoId
			"convAutoId": m.ConvAutoId,
			// 更新 createTime
			"createTime": m.CreateTime,
			// 更新 contentType
			"contentType": m.ContentType,
			// 更新 content
			"content": m.Content,
			// 更新 title
			"title": m.Title,
			// 更新 ext
			"ext": m.Ext,
			// 更新 options
			"options": m.Options,
			// 更新 convId
			"convId": m.ConvId,
			// 更新 userId
			"userId": m.UserId,
		}),
	}).Create(m).Error

}

func BatchInsert(tx *gorm.DB, notices []*Notice, rc *redis.Redis) error {
	ctx := context.Background()
	for _, m := range notices {
		if m.ConvAutoId == 0 {
			// 获取maxConvAutoId
			convAutoId, err := GetMaxConvAutoId(ctx, rc, m.ConvId, 1)
			if err != nil {
				return err
			}
			m.ConvAutoId = convAutoId
		}
		if m.NoticeId == "" {
			m.NoticeId = pb.ServerNoticeId(m.ConvId, m.ConvAutoId, m.UserId)
		}
		if m.CreateTime == 0 {
			m.CreateTime = time.Now().UnixMilli()
		}
		if m.Ext == nil {
			m.Ext = make([]byte, 0)
		}
		if m.Content == nil {
			m.Content = make([]byte, 0)
		}
	}
	// 忽略唯一索引冲突
	return tx.Model(&Notice{}).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(notices).Error
}

func GetMaxConvAutoId(ctx context.Context, rc *redis.Redis, convId string, incr int64) (int64, error) {
	logger := logx.WithContext(ctx)
	// incr redis
	key := rediskey.NoticeConvAutoId(convId)
	maxHKey := convId
	val, err := rc.HincrbyCtx(ctx, key, maxHKey, int(incr))
	if err != nil {
		logger.Errorf("incr redis err: %v", err)
		return 0, err
	}
	return int64(val), nil
}

func GetMinConvAutoId(ctx context.Context, rc *redis.Redis, convId string, userId string, deviceId string) (int64, error) {
	logger := logx.WithContext(ctx)
	key := rediskey.NoticeConvAutoId(convId)
	hkey := fmt.Sprintf("%s:%s", userId, deviceId)
	val, err := rc.HgetCtx(ctx, key, hkey)
	if err == redis.Nil || val == "" {
		// 设置为maxConvAutoId
		maxConvAutoId, err := GetMaxConvAutoId(ctx, rc, convId, 0)
		if err != nil {
			return 0, err
		}
		err = rc.HsetCtx(ctx, key, hkey, strconv.FormatInt(maxConvAutoId-1, 10))
		if err != nil {
			logger.Errorf("set minConvAutoId err: %v", err)
			return 0, err
		}
		return maxConvAutoId - 1, nil
	}
	if err != nil {
		logger.Errorf("get minConvAutoId err: %v", err)
		return 0, err
	}
	minConvAutoId, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		logger.Errorf("parse minConvAutoId err: %v", err)
		return 0, err
	}
	return minConvAutoId, nil
}

func SetMinConvAutoId(ctx context.Context, rc *redis.Redis, convId string, userId string, deviceId string, seq int64) error {
	// hset
	key := rediskey.NoticeConvAutoId(convId)
	hkey := fmt.Sprintf("%s:%s", userId, deviceId)
	return rc.HsetCtx(ctx, key, hkey, strconv.FormatInt(seq, 10))
}

func GetNotice(ctx context.Context, tx *gorm.DB, rc *redis.Redis, convId string, userRedisKey string, userId string, deviceId string, minSeq int64, maxSeq int64) (*Notice, error) {
	logger := logx.WithContext(ctx)
	sortSetKey := rediskey.NoticeSortSetKey(convId, userRedisKey, userId, deviceId)
	// get最小的 ZRANGEBYSCORE key -inf +inf LIMIT 0 1
	pairs, err := rc.ZrangebyscoreWithScoresAndLimitCtx(ctx, sortSetKey, -405, 0, 0, 1)
	if err != nil {
		// 如果是 redis.Nil 则表示没有数据
		if err == redis.Nil {
			return getNoticeFromMysql(ctx, tx, rc, convId, userRedisKey, userId, deviceId, minSeq, maxSeq, sortSetKey)
		} else {
			logger.Errorf("zpopmin err: %v", err)
			return getNoticeFromMysql(ctx, tx, rc, convId, userRedisKey, userId, deviceId, minSeq, maxSeq, sortSetKey)
		}
	}
	if len(pairs) == 0 {
		// 没有数据
		return getNoticeFromMysql(ctx, tx, rc, convId, userRedisKey, userId, deviceId, minSeq, maxSeq, sortSetKey)
	}
	noticeId := pairs[0].Key
	// 如果 noticeId == xredis.NotFound 则表示真的没有数据 直接返回
	// 有bug 注掉
	//if noticeId == xredis.NotFound {
	//	return nil, nil
	//}
	// 使用noticeId查询notice
	var notice *Notice
	notice, err = GetNoticeById(ctx, tx, rc, noticeId)
	if err != nil {
		logger.Errorf("find notice err: %v", err)
		return nil, err
	}
	return notice, nil
}

func getNoticeFromMysql(ctx context.Context, tx *gorm.DB, rc *redis.Redis, convId string, keyId string, userId string, deviceId string, minSeq int64, maxSeq int64, sortSetKey string) (*Notice, error) {
	logger := logx.WithContext(ctx)
	// 先删掉redis中的数据
	err := flushNoticeZSet(ctx, rc, convId, keyId, userId, deviceId)
	// 直接查询mysql
	var notices []*Notice
	err = tx.Model(&Notice{}).
		Where(""+
			"(convId = ? and userId = ?)"+
			" and convAutoId > ? and convAutoId <= ?",
			convId, userId,
			minSeq, maxSeq).
		Order("convAutoId asc").
		Limit(1000).
		Find(&notices).Error
	if err != nil {
		logger.Errorf("find notice err: %v", err)
		return nil, err
	}
	if len(notices) == 0 {
		// redis插入一条不存在的数据，防止缓存穿透
		// 有bug 注掉
		//err := xredis.ZAddsEx(rc, ctx, sortSetKey, rediskey.NoticeSortSetExpire(), -1, xredis.NotFound)
		//if err != nil {
		//	logger.Errorf("redis zadd err: %v", err)
		//}
		return nil, nil
	}
	// 保存到redis
	args := make([]interface{}, 0, len(notices)*2)
	for _, notice := range notices {
		args = append(args, notice.ConvAutoId, notice.NoticeId)
	}
	err = xredis.ZAddsEx(rc, ctx, sortSetKey, rediskey.NoticeSortSetExpire(), args...)
	if err != nil {
		logger.Errorf("redis zadd err: %v", err)
	}
	return notices[0], nil
}

func flushNoticeZSet(ctx context.Context, rc *redis.Redis, convId string, keyId string, userId string, deviceId string) error {
	logger := logx.WithContext(ctx)
	// 先删掉redis中的数据
	sortSetKey := rediskey.NoticeSortSetKey(convId, keyId, userId, deviceId)
	_, err := rc.DelCtx(ctx, sortSetKey)
	if err != nil {
		logger.Errorf("redis del err: %v", err)
	}
	return err
}

func DelNoticeZSet(ctx context.Context, rc *redis.Redis, convId string, keyId string, userId string, deviceId string, seq int64) error {
	logger := logx.WithContext(ctx)
	sortSetKey := rediskey.NoticeSortSetKey(convId, keyId, userId, deviceId)
	// zrembyseq
	_, err := rc.ZremrangebyscoreCtx(ctx, sortSetKey, seq, seq)
	if err != nil {
		logger.Errorf("redis zrembyseq err: %v", err)
	}
	return err
}

func GetNoticeById(ctx context.Context, tx *gorm.DB, rc *redis.Redis, id string) (*Notice, error) {
	logger := logx.WithContext(ctx)
	// 先从redis查询
	stringKey := rediskey.NoticeStringKey(id)
	val, err := rc.GetCtx(ctx, stringKey)
	if err != nil {
		// 如果是 redis.Nil 则表示没有数据 需要从mysql查询
		if err == redis.Nil {
			return getNoticeByIdFromMysql(ctx, tx, rc, id)
		} else {
			logger.Errorf("redis get err: %v", err)
			return getNoticeByIdFromMysql(ctx, tx, rc, id)
		}
	}
	// 如果 val == xredis.NotFound 则表示真的没有数据 直接返回
	if val == xredis.NotFound {
		return nil, nil
	}
	// 反序列化
	notice := &Notice{}
	err = json.Unmarshal([]byte(val), notice)
	if err != nil {
		logger.Errorf("json unmarshal err: %v", err)
		return getNoticeByIdFromMysql(ctx, tx, rc, id)
	}
	return notice, nil
}

func getNoticeByIdFromMysql(ctx context.Context, tx *gorm.DB, rc *redis.Redis, id string) (*Notice, error) {
	logger := logx.WithContext(ctx)
	// 先删掉redis中的数据
	err := FlushNoticeString(ctx, rc, id)
	// 直接查询mysql
	var notice Notice
	err = tx.Model(&Notice{}).
		Where("noticeId = ?", id).Limit(1).
		Find(&notice).Error
	if err != nil {
		logger.Errorf("find notice err: %v", err)
		return nil, err
	}
	// 保存到redis
	err = rc.SetexCtx(ctx, rediskey.NoticeStringKey(id), utils.AnyToString(notice), rediskey.NoticeStringExpire())
	if err != nil {
		logger.Errorf("save notice to redis err: %v", err)
	}
	return &notice, nil
}

func FlushNoticeString(ctx context.Context, rc *redis.Redis, id string) error {
	logger := logx.WithContext(ctx)
	// 先删掉redis中的数据
	stringKey := rediskey.NoticeStringKey(id)
	_, err := rc.DelCtx(ctx, stringKey)
	if err != nil {
		logger.Errorf("redis del err: %v", err)
	}
	return err
}

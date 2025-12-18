package data

import (
	"encoding/json"
	"errors"
	"jd-go-stock/backend/db"
	"jd-go-stock/backend/logger"
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	TushareToken           string `json:"tushareToken"`
	LocalPushEnable        bool   `json:"localPushEnable"`
	DingPushEnable         bool   `json:"dingPushEnable"`
	DingRobot              string `json:"dingRobot"`
	UpdateBasicInfoOnStart bool   `json:"updateBasicInfoOnStart"`
	RefreshInterval        int64  `json:"refreshInterval"`
	OpenAiEnable           bool   `json:"openAiEnable"`
	Prompt                 string `json:"prompt"`
	CheckUpdate            bool   `json:"checkUpdate"`
	QuestionTemplate       string `json:"questionTemplate"`
	CrawlTimeOut           int64  `json:"crawlTimeOut"`
	KDays                  int64  `json:"kDays"`
	EnableDanmu            bool   `json:"enableDanmu"`
	BrowserPath            string `json:"browserPath"`
	EnableNews             bool   `json:"enableNews"`
	DarkTheme              bool   `json:"darkTheme"`
	BrowserPoolSize        int    `json:"browserPoolSize"`
	EnableFund             bool   `json:"enableFund"`
	EnablePushNews         bool   `json:"enablePushNews"`
	EnableOnlyPushRedNews  bool   `json:"enableOnlyPushRedNews"`
	SponsorCode            string `json:"sponsorCode"`
	HttpProxy              string `json:"httpProxy"`
	HttpProxyEnabled       bool   `json:"httpProxyEnabled"`
	EnableAgent            bool   `json:"enableAgent"`
	QgqpBId                string `json:"qgqpBId" gorm:"column:qgqp_b_id"`
}

func (receiver Settings) TableName() string {
	return "settings"
}

type AIConfig struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string  `json:"name"`
	BaseUrl     string  `json:"baseUrl"`
	ApiKey      string  `json:"apiKey" `
	ModelName   string  `json:"modelName"`
	MaxTokens   int     `json:"maxTokens"`
	Temperature float64 `json:"temperature"`
	TimeOut     int     `json:"timeOut"`
}

func (AIConfig) TableName() string {
	return "ai_config"
}

type SettingConfig struct {
	*Settings
	AiConfigs []*AIConfig `json:"aiConfigs"`
}

type SettingsApi struct {
	Config *SettingConfig
}

func NewSettingsApi() *SettingsApi {
	return &SettingsApi{
		Config: GetSettingConfig(),
	}
}

func (s *SettingsApi) Export() string {
	d, _ := json.MarshalIndent(s.Config, "", "    ")
	return string(d)
}

func UpdateConfig(s *SettingConfig) string {
	count := int64(0)
	db.Dao.Model(&Settings{}).Count(&count)
	if count > 0 {
		db.Dao.Model(&Settings{}).Where("id=?", s.ID).Updates(map[string]any{
			"local_push_enable":          s.LocalPushEnable,
			"ding_push_enable":           s.DingPushEnable,
			"ding_robot":                 s.DingRobot,
			"update_basic_info_on_start": s.UpdateBasicInfoOnStart,
			"refresh_interval":           s.RefreshInterval,
			"open_ai_enable":             s.OpenAiEnable,
			"tushare_token":              s.TushareToken,
			"prompt":                     s.Prompt,
			"check_update":               s.CheckUpdate,
			"question_template":          s.QuestionTemplate,
			"crawl_time_out":             s.CrawlTimeOut,
			"k_days":                     s.KDays,
			"enable_danmu":               s.EnableDanmu,
			"browser_path":               s.BrowserPath,
			"enable_news":                s.EnableNews,
			"dark_theme":                 s.DarkTheme,
			"enable_fund":                s.EnableFund,
			"enable_push_news":           s.EnablePushNews,
			"enable_only_push_red_news":  s.EnableOnlyPushRedNews,
			"sponsor_code":               s.SponsorCode,
			"http_proxy":                 s.HttpProxy,
			"http_proxy_enabled":         s.HttpProxyEnabled,
			"enable_agent":               s.EnableAgent,
			"qgqp_b_id":                  s.QgqpBId,
		})

		//更新AiConfig
		err := updateAiConfigs(s.AiConfigs)
		if err != nil {
			logger.SugaredLogger.Errorf("更新AI模型服务配置失败: %v", err)
			return "更新AI模型服务配置失败: " + err.Error()
		}
	} else {
		logger.SugaredLogger.Infof("未找到配置，创建默认配置")
		// 创建主配置
		result := db.Dao.Model(&Settings{}).Create(&Settings{})
		if result.Error != nil {
			logger.SugaredLogger.Error("创建配置失败:", result.Error)
			return "创建配置失败: " + result.Error.Error()
		}
	}
	return "保存成功！"
}

func updateAiConfigs(aiConfigs []*AIConfig) error {
	if len(aiConfigs) == 0 {
		err := db.Dao.Exec("DELETE FROM ai_config").Error
		if err != nil {
			return err
		}
		return db.Dao.Exec("DELETE FROM sqlite_sequence WHERE name='ai_config'").Error
	}
	var ids []uint
	lo.ForEach(aiConfigs, func(item *AIConfig, index int) {
		ids = append(ids, item.ID)
	})
	var existAiConfigs []*AIConfig
	err := db.Dao.Model(&AIConfig{}).Select("id").Where("id in (?) ", ids).Find(&existAiConfigs).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	idMap := make(map[uint]bool)
	lo.ForEach(existAiConfigs, func(item *AIConfig, index int) {
		idMap[item.ID] = true
	})
	var addAiConfigs []*AIConfig
	var notDeleteIds []uint
	var e error
	lo.ForEach(aiConfigs, func(item *AIConfig, index int) {
		if e != nil {
			return
		}
		if !idMap[item.ID] {
			addAiConfigs = append(addAiConfigs, item)
		} else {
			notDeleteIds = append(notDeleteIds, item.ID)
			e = db.Dao.Model(&AIConfig{}).Where("id=?", item.ID).Updates(map[string]interface{}{
				"name":        item.Name,
				"base_url":    item.BaseUrl,
				"api_key":     item.ApiKey,
				"model_name":  item.ModelName,
				"max_tokens":  item.MaxTokens,
				"temperature": item.Temperature,
				"time_out":    item.TimeOut,
			}).Error
			if e != nil {
				return
			}
		}
	})
	if e != nil {
		return e
	}
	//删除旧的配置
	if len(notDeleteIds) > 0 {
		err = db.Dao.Exec("DELETE FROM ai_config WHERE id NOT IN ?", notDeleteIds).Error
		if err != nil {
			return err
		}
	}
	logger.SugaredLogger.Infof("更新aiConfigs +%d", len(addAiConfigs))
	//批量新增的配置
	err = db.Dao.CreateInBatches(addAiConfigs, len(addAiConfigs)).Error
	return err
}

func GetSettingConfig() *SettingConfig {
	settingConfig := &SettingConfig{}
	settings := &Settings{}
	aiConfigs := make([]*AIConfig, 0)
	// 处理数据库查询可能返回的空结果
	result := db.Dao.Model(&Settings{}).First(settings)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 初始化默认设置并保存到数据库
		settings = &Settings{OpenAiEnable: false, CrawlTimeOut: 60}
		db.Dao.Create(settings)
	}

	if settings.OpenAiEnable {
		// 处理AI配置查询可能出现的错误
		result = db.Dao.Model(&AIConfig{}).Find(&aiConfigs)
		if result.Error != nil {
			logger.SugaredLogger.Error("查询AI配置失败:", result.Error)
		} else if len(aiConfigs) > 0 {
			lo.ForEach(aiConfigs, func(item *AIConfig, index int) {
				if item.TimeOut <= 0 {
					item.TimeOut = 60 * 5
				}
			})
		}
		if settings.CrawlTimeOut <= 0 {
			settings.CrawlTimeOut = 60
		}
		if settings.KDays < 30 {
			settings.KDays = 120
		}
	}
	if settings.BrowserPath == "" {
		settings.BrowserPath, _ = CheckBrowser()
	}
	if settings.BrowserPoolSize <= 0 {
		settings.BrowserPoolSize = 1
	}
	settingConfig.Settings = settings
	settingConfig.AiConfigs = aiConfigs

	return settingConfig
}

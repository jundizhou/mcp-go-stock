package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// @Author spark
// @Date 2025/2/6 15:25
// @Desc
//-----------------------------------------------------------------------------------

type GitHubReleaseVersion struct {
	Url       string `json:"url"`
	AssetsUrl string `json:"assets_url"`
	UploadUrl string `json:"upload_url"`
	HtmlUrl   string `json:"html_url"`
	Id        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		UserViewType      string `json:"user_view_type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeId          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		Url      string `json:"url"`
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			UserViewType      string `json:"user_view_type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadUrl string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballUrl string `json:"tarball_url"`
	ZipballUrl string `json:"zipball_url"`
	Body       string `json:"body"`
	Tag        Tag    `json:"tag"`
	Commit     Commit `json:"commit"`
}

type Tag struct {
	Ref    string `json:"ref"`
	NodeId string `json:"node_id"`
	Url    string `json:"url"`
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"object"`
}

type Commit struct {
	Sha     string `json:"sha"`
	NodeId  string `json:"node_id"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Author  struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"author"`
	Committer struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"committer"`
	Tree struct {
		Sha string `json:"sha"`
		Url string `json:"url"`
	} `json:"tree"`
	Message string `json:"message"`
	Parents []struct {
		Sha     string `json:"sha"`
		Url     string `json:"url"`
		HtmlUrl string `json:"html_url"`
	} `json:"parents"`
	Verification struct {
		Verified   bool        `json:"verified"`
		Reason     string      `json:"reason"`
		Signature  interface{} `json:"signature"`
		Payload    interface{} `json:"payload"`
		VerifiedAt interface{} `json:"verified_at"`
	} `json:"verification"`
}

type AIResponseResult struct {
	gorm.Model
	ChatId    string                `json:"chatId"`
	ModelName string                `json:"modelName"`
	StockCode string                `json:"stockCode"`
	StockName string                `json:"stockName"`
	Question  string                `json:"question"`
	Content   string                `json:"content"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

func (receiver AIResponseResult) TableName() string {
	return "ai_response_result"
}

type VersionInfo struct {
	gorm.Model
	Version           string                `json:"version"`
	Content           string                `json:"content"`
	Icon              string                `json:"icon"`
	Alipay            string                `json:"alipay"`
	Wxpay             string                `json:"wxpay"`
	Wxgzh             string                `json:"wxgzh"`
	BuildTimeStamp    int64                 `json:"buildTimeStamp"`
	OfficialStatement string                `json:"officialStatement"`
	IsDel             soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

func (receiver VersionInfo) TableName() string {
	return "version_info"
}

type StockInfoHK struct {
	gorm.Model
	Code     string                `json:"code"`
	Name     string                `json:"name"`
	FullName string                `json:"fullName"`
	EName    string                `json:"eName"`
	IsDel    soft_delete.DeletedAt `gorm:"softDelete:flag"`
	BKName   string                `json:"bk_name"`
	BKCode   string                `json:"bk_code"`
}

func (receiver StockInfoHK) TableName() string {
	return "stock_base_info_hk"
}

type StockInfoUS struct {
	gorm.Model
	Code     string                `json:"code"`
	Name     string                `json:"name"`
	FullName string                `json:"fullName"`
	EName    string                `json:"eName"`
	Exchange string                `json:"exchange"`
	Type     string                `json:"type"`
	IsDel    soft_delete.DeletedAt `gorm:"softDelete:flag"`
	BKName   string                `json:"bk_name"`
	BKCode   string                `json:"bk_code"`
}

func (receiver StockInfoUS) TableName() string {
	return "stock_base_info_us"
}

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Param   string `json:"param"`
		Type    string `json:"type"`
	} `json:"error"`
}

type PromptTemplate struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
	Content   string `json:"content"`
	Type      string `json:"type"`
}

func (p PromptTemplate) TableName() string {
	return "prompt_templates"
}

type Prompt struct {
	ID      int    `json:"ID"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type Telegraph struct {
	gorm.Model
	Time            string          `json:"time"`
	DataTime        *time.Time      `json:"dataTime" gorm:"index"`
	Title           string          `json:"title" gorm:"index"`
	Content         string          `json:"content" gorm:"index"`
	SubjectTags     []string        `json:"subjects" gorm:"-:all"`
	StocksTags      []string        `json:"stocks" gorm:"-:all"`
	IsRed           bool            `json:"isRed" gorm:"index"`
	Url             string          `json:"url"`
	Source          string          `json:"source" gorm:"index"`
	TelegraphTags   []TelegraphTags `json:"tags" gorm:"-:migration;foreignKey:TelegraphId"`
	SentimentResult string          `json:"sentimentResult" gorm:"index"`
}
type TelegraphTags struct {
	gorm.Model
	TagId       uint `json:"tagId"`
	TelegraphId uint `json:"telegraphId"`
}

func (t TelegraphTags) TableName() string {
	return "telegraph_tags"
}

type Tags struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}

func (p Tags) TableName() string {
	return "tags"
}

func (p Telegraph) TableName() string {
	return "telegraph_list"
}

type SinaStockInfo struct {
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Engname       string `json:"engname"`
	Tradetype     string `json:"tradetype"`
	Lasttrade     string `json:"lasttrade"`
	Prevclose     string `json:"prevclose"`
	Open          string `json:"open"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Volume        string `json:"volume"`
	Currentvolume string `json:"currentvolume"`
	Amount        string `json:"amount"`
	Ticktime      string `json:"ticktime"`
	Buy           string `json:"buy"`
	Sell          string `json:"sell"`
	High52Week    string `json:"high_52week"`
	Low52Week     string `json:"low_52week"`
	Eps           string `json:"eps"`
	Dividend      string `json:"dividend"`
	StocksSum     string `json:"stocks_sum"`
	Pricechange   string `json:"pricechange"`
	Changepercent string `json:"changepercent"`
	MarketValue   string `json:"market_value"`
	PeRatio       string `json:"pe_ratio"`
}

type LongTigerRankData struct {
	ACCUMAMOUNT      float64 `json:"ACCUM_AMOUNT"`
	BILLBOARDBUYAMT  float64 `json:"BILLBOARD_BUY_AMT"`
	BILLBOARDDEALAMT float64 `json:"BILLBOARD_DEAL_AMT"`
	BILLBOARDNETAMT  float64 `json:"BILLBOARD_NET_AMT"`
	BILLBOARDSELLAMT float64 `json:"BILLBOARD_SELL_AMT"`
	CHANGERATE       float64 `json:"CHANGE_RATE"`
	CLOSEPRICE       float64 `json:"CLOSE_PRICE"`
	DEALAMOUNTRATIO  float64 `json:"DEAL_AMOUNT_RATIO"`
	DEALNETRATIO     float64 `json:"DEAL_NET_RATIO"`
	EXPLAIN          string  `json:"EXPLAIN"`
	EXPLANATION      string  `json:"EXPLANATION"`
	FREEMARKETCAP    float64 `json:"FREE_MARKET_CAP"`
	SECUCODE         string  `json:"SECUCODE" gorm:"index"`
	SECURITYCODE     string  `json:"SECURITY_CODE"`
	SECURITYNAMEABBR string  `json:"SECURITY_NAME_ABBR"`
	SECURITYTYPECODE string  `json:"SECURITY_TYPE_CODE"`
	TRADEDATE        string  `json:"TRADE_DATE" gorm:"index"`
	TURNOVERRATE     float64 `json:"TURNOVERRATE"`
}

func (l LongTigerRankData) TableName() string {
	return "long_tiger_rank"
}

type TVNews struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Published  int    `json:"published"`
	Urgency    int    `json:"urgency"`
	Permission string `json:"permission"`
	StoryPath  string `json:"storyPath"`
	Provider   struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		LogoId string `json:"logo_id"`
	} `json:"provider"`
}
type TVNewsDetail struct {
	ShortDescription string `json:"shortDescription"`
	Tags             []struct {
		Title string `json:"title"`
		Args  []struct {
			Id    string `json:"id"`
			Value string `json:"value"`
		} `json:"args"`
	} `json:"tags"`
	Copyright string `json:"copyright"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Published int    `json:"published"`
	Urgency   int    `json:"urgency"`
	StoryPath string `json:"storyPath"`
}

type XUEQIUHot struct {
	Data struct {
		Items     []HotItem `json:"items"`
		ItemsSize int       `json:"items_size"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type HotItem struct {
	Type         int         `json:"type"`
	Code         string      `json:"code"`
	Name         string      `json:"name"`
	Value        float64     `json:"value"`
	Increment    int         `json:"increment"`
	RankChange   int         `json:"rank_change"`
	HasExist     interface{} `json:"has_exist"`
	Symbol       string      `json:"symbol"`
	Percent      float64     `json:"percent"`
	Current      float64     `json:"current"`
	Chg          float64     `json:"chg"`
	Exchange     string      `json:"exchange"`
	StockType    int         `json:"stock_type"`
	SubType      string      `json:"sub_type"`
	Ad           int         `json:"ad"`
	AdId         interface{} `json:"ad_id"`
	ContentId    interface{} `json:"content_id"`
	Page         interface{} `json:"page"`
	Model        interface{} `json:"model"`
	Location     interface{} `json:"location"`
	TradeSession interface{} `json:"trade_session"`
	CurrentExt   interface{} `json:"current_ext"`
	PercentExt   interface{} `json:"percent_ext"`
}

type HotEvent struct {
	PicSize     interface{} `json:"pic_size"`
	Tag         string      `json:"tag"`
	Id          int         `json:"id"`
	Pic         string      `json:"pic"`
	Hot         int         `json:"hot"`
	StatusCount int         `json:"status_count"`
	Content     string      `json:"content"`
}

type GDP struct {
	REPORTDATE           string  `json:"REPORT_DATE" md:"报告时间"`
	TIME                 string  `json:"TIME" md:"报告期"`
	DOMESTICLPRODUCTBASE float64 `json:"DOMESTICL_PRODUCT_BASE" md:"国内生产总值(亿元)"`
	SUMSAME              float64 `json:"SUM_SAME" md:"国内生产总值同比增长(%)"`
	FIRSTPRODUCTBASE     float64 `json:"FIRST_PRODUCT_BASE" md:"第一产业(亿元)"`
	FIRSTSAME            int     `json:"FIRST_SAME" md:"第一产业同比增长(%)"`
	SECONDPRODUCTBASE    float64 `json:"SECOND_PRODUCT_BASE" md:"第二产业(亿元)"`
	SECONDSAME           float64 `json:"SECOND_SAME" md:"第二产业同比增长(%)"`
	THIRDPRODUCTBASE     float64 `json:"THIRD_PRODUCT_BASE" md:"第三产业(亿元)"`
	THIRDSAME            float64 `json:"THIRD_SAME" md:"第三产业同比增长(%)"`
}
type CPI struct {
	REPORTDATE         string  `json:"REPORT_DATE" md:"报告时间"`
	TIME               string  `json:"TIME" md:"报告期"`
	NATIONALBASE       float64 `json:"NATIONAL_BASE" md:"全国当月"`
	NATIONALSAME       float64 `json:"NATIONAL_SAME" md:"全国当月同比增长(%)"`
	NATIONALSEQUENTIAL float64 `json:"NATIONAL_SEQUENTIAL" md:"全国当月环比增长(%)"`
	NATIONALACCUMULATE float64 `json:"NATIONAL_ACCUMULATE" md:"全国当月累计"`
	CITYBASE           float64 `json:"CITY_BASE" md:"城市当月"`
	CITYSAME           float64 `json:"CITY_SAME" md:"城市当月同比增长(%)"`
	CITYSEQUENTIAL     float64 `json:"CITY_SEQUENTIAL" md:"城市当月环比增长(%)"`
	CITYACCUMULATE     int     `json:"CITY_ACCUMULATE" md:"城市当月累计"`
	RURALBASE          float64 `json:"RURAL_BASE" md:"农村当月"`
	RURALSAME          float64 `json:"RURAL_SAME" md:"农村当月同比增长(%)"`
	RURALSEQUENTIAL    int     `json:"RURAL_SEQUENTIAL" md:"农村当月环比增长(%)"`
	RURALACCUMULATE    float64 `json:"RURAL_ACCUMULATE" md:"农村当月累计"`
}
type PPI struct {
	REPORTDATE     string  `json:"REPORT_DATE" md:"报告时间"`
	TIME           string  `json:"TIME" md:"报告期"`
	BASE           float64 `json:"BASE" md:"当月"`
	BASESAME       float64 `json:"BASE_SAME" md:"当月同比增长(%)"`
	BASEACCUMULATE float64 `json:"BASE_ACCUMULATE" md:"累计"`
}
type PMI struct {
	REPORTDATE string  `md:"报告时间" json:"REPORT_DATE"`
	TIME       string  `md:"报告期" json:"TIME"`
	MAKEINDEX  float64 `md:"制造业指数" json:"MAKE_INDEX"`
	MAKESAME   float64 `md:"制造业指数同比增长(%)" json:"MAKE_SAME"`
	NMAKEINDEX float64 `md:"非制造业" json:"NMAKE_INDEX"`
	NMAKESAME  float64 `md:"非制造业同比增长(%)" json:"NMAKE_SAME"`
}

type DCResp struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type GDPResult struct {
	Pages int   `json:"pages"`
	Data  []GDP `json:"data"`
	Count int   `json:"count"`
}
type CPIResult struct {
	Pages int   `json:"pages"`
	Data  []CPI `json:"data"`
	Count int   `json:"count"`
}

type PPIResult struct {
	Pages int   `json:"pages"`
	Data  []PPI `json:"data"`
	Count int   `json:"count"`
}
type PMIResult struct {
	Pages int   `json:"pages"`
	Data  []PMI `json:"data"`
	Count int   `json:"count"`
}
type GDPResp struct {
	DCResp
	GDPResult GDPResult `json:"result"`
}

type CPIResp struct {
	DCResp
	CPIResult CPIResult `json:"result"`
}

type PPIResp struct {
	DCResp
	PPIResult PPIResult `json:"result"`
}
type PMIResp struct {
	DCResp
	PMIResult PMIResult `json:"result"`
}

type OldSettings struct {
	gorm.Model
	TushareToken           string `json:"tushareToken"`
	LocalPushEnable        bool   `json:"localPushEnable"`
	DingPushEnable         bool   `json:"dingPushEnable"`
	DingRobot              string `json:"dingRobot"`
	UpdateBasicInfoOnStart bool   `json:"updateBasicInfoOnStart"`
	RefreshInterval        int64  `json:"refreshInterval"`

	OpenAiEnable      bool    `json:"openAiEnable"`
	OpenAiBaseUrl     string  `json:"openAiBaseUrl"`
	OpenAiApiKey      string  `json:"openAiApiKey"`
	OpenAiModelName   string  `json:"openAiModelName"`
	OpenAiMaxTokens   int     `json:"openAiMaxTokens"`
	OpenAiTemperature float64 `json:"openAiTemperature"`
	OpenAiApiTimeOut  int     `json:"openAiApiTimeOut"`
	Prompt            string  `json:"prompt"`
	CheckUpdate       bool    `json:"checkUpdate"`
	QuestionTemplate  string  `json:"questionTemplate"`
	CrawlTimeOut      int64   `json:"crawlTimeOut"`
	KDays             int64   `json:"kDays"`
	EnableDanmu       bool    `json:"enableDanmu"`
	BrowserPath       string  `json:"browserPath"`
	EnableNews        bool    `json:"enableNews"`
	DarkTheme         bool    `json:"darkTheme"`
	BrowserPoolSize   int     `json:"browserPoolSize"`
	EnableFund        bool    `json:"enableFund"`
	EnablePushNews    bool    `json:"enablePushNews"`
	SponsorCode       string  `json:"sponsorCode"`
}

func (receiver OldSettings) TableName() string {
	return "settings"
}

type ReutersNews struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Result     struct {
		ParentSectionName string `json:"parent_section_name"`
		Pagination        struct {
			Size         int    `json:"size"`
			ExpectedSize int    `json:"expected_size"`
			TotalSize    int    `json:"total_size"`
			Orderby      string `json:"orderby"`
		} `json:"pagination"`
		DateModified time.Time `json:"date_modified"`
		FetchType    string    `json:"fetch_type"`
		Articles     []struct {
			Id                          string    `json:"id"`
			CanonicalUrl                string    `json:"canonical_url"`
			Website                     string    `json:"website"`
			Web                         string    `json:"web"`
			Native                      string    `json:"native"`
			UpdatedTime                 time.Time `json:"updated_time"`
			PublishedTime               time.Time `json:"published_time"`
			ArticleType                 string    `json:"article_type"`
			DisplayMyNews               bool      `json:"display_my_news"`
			DisplayNewsletterSignup     bool      `json:"display_newsletter_signup"`
			DisplayNotifications        bool      `json:"display_notifications"`
			DisplayRelatedMedia         bool      `json:"display_related_media"`
			DisplayRelatedOrganizations bool      `json:"display_related_organizations"`
			ContentCode                 string    `json:"content_code"`
			Source                      struct {
				Name         string `json:"name"`
				OriginalName string `json:"original_name"`
			} `json:"source"`
			Title            string `json:"title"`
			BasicHeadline    string `json:"basic_headline"`
			Distributor      string `json:"distributor"`
			Description      string `json:"description"`
			PrimaryMediaType string `json:"primary_media_type,omitempty"`
			PrimaryTag       struct {
				ShortBio    string `json:"short_bio"`
				Description string `json:"description"`
				Slug        string `json:"slug"`
				Text        string `json:"text"`
				TopicUrl    string `json:"topic_url"`
				CanFollow   bool   `json:"can_follow,omitempty"`
				IsTopic     bool   `json:"is_topic,omitempty"`
			} `json:"primary_tag"`
			WordCount   int `json:"word_count"`
			ReadMinutes int `json:"read_minutes"`
			Kicker      struct {
				Path  string   `json:"path"`
				Names []string `json:"names"`
				Name  string   `json:"name,omitempty"`
			} `json:"kicker"`
			AdTopics  []string `json:"ad_topics"`
			Thumbnail struct {
				Url                   string    `json:"url"`
				Caption               string    `json:"caption,omitempty"`
				Type                  string    `json:"type"`
				ResizerUrl            string    `json:"resizer_url"`
				Location              string    `json:"location,omitempty"`
				Id                    string    `json:"id"`
				Authors               string    `json:"authors,omitempty"`
				AltText               string    `json:"alt_text"`
				Width                 int       `json:"width"`
				Height                int       `json:"height"`
				Subtitle              string    `json:"subtitle"`
				Slug                  string    `json:"slug,omitempty"`
				UpdatedAt             time.Time `json:"updated_at"`
				Company               string    `json:"company,omitempty"`
				PurchaseLicensingPath string    `json:"purchase_licensing_path,omitempty"`
			} `json:"thumbnail"`
			Authors []struct {
				Id        string `json:"id,omitempty"`
				Name      string `json:"name"`
				FirstName string `json:"first_name,omitempty"`
				LastName  string `json:"last_name,omitempty"`
				Company   string `json:"company"`
				Thumbnail struct {
					Url        string `json:"url"`
					Type       string `json:"type"`
					ResizerUrl string `json:"resizer_url"`
				} `json:"thumbnail"`
				SocialLinks []struct {
					Site string `json:"site"`
					Url  string `json:"url"`
				} `json:"social_links,omitempty"`
				Byline      string `json:"byline"`
				Description string `json:"description,omitempty"`
				TopicUrl    string `json:"topic_url,omitempty"`
				Role        string `json:"role,omitempty"`
			} `json:"authors"`
			DisplayTime   time.Time `json:"display_time"`
			ThumbnailDark struct {
				Url        string    `json:"url"`
				Type       string    `json:"type"`
				ResizerUrl string    `json:"resizer_url"`
				Id         string    `json:"id"`
				AltText    string    `json:"alt_text"`
				Width      int       `json:"width"`
				Height     int       `json:"height"`
				Subtitle   string    `json:"subtitle"`
				UpdatedAt  time.Time `json:"updated_at"`
			} `json:"thumbnail_dark,omitempty"`
		} `json:"articles"`
		Section struct {
			Id          string `json:"id"`
			AdUnitCode  string `json:"ad_unit_code"`
			Website     string `json:"website"`
			Name        string `json:"name"`
			PageTitle   string `json:"page_title"`
			CanFollow   bool   `json:"can_follow"`
			Language    string `json:"language"`
			Type        string `json:"type"`
			Advertising struct {
				Sponsored string `json:"sponsored"`
			} `json:"advertising"`
			VideoPlaylistId  string `json:"video_playlistId"`
			MobileAdUnitPath string `json:"mobile_ad_unit_path"`
			AdUnitPath       string `json:"ad_unit_path"`
			CollectionAlias  string `json:"collection_alias"`
			SectionAbout     string `json:"section_about"`
			Title            string `json:"title"`
			Personalization  struct {
				Id        string `json:"id"`
				Type      string `json:"type"`
				ShowTags  bool   `json:"show_tags"`
				CanFollow bool   `json:"can_follow"`
			} `json:"personalization"`
		} `json:"section"`
		AdUnitPath   string `json:"ad_unit_path"`
		ResponseTime int64  `json:"response_time"`
	} `json:"result"`
	Id string `json:"_id"`
}

type InteractiveAnswer struct {
	PageNo      int                        `json:"pageNo"`
	PageSize    int                        `json:"pageSize"`
	TotalRecord int                        `json:"totalRecord"`
	TotalPage   int                        `json:"totalPage"`
	Results     []InteractiveAnswerResults `json:"results"`
	Count       bool                       `json:"count"`
}

type InteractiveAnswerResults struct {
	EsId             string   `json:"esId" md:"-"`
	IndexId          string   `json:"indexId" md:"-"`
	ContentType      int      `json:"contentType" md:"-"`
	Trade            []string `json:"trade"  md:"行业名称"`
	MainContent      string   `json:"mainContent" md:"投资者提问"`
	StockCode        string   `json:"stockCode" md:"股票代码"`
	Secid            string   `json:"secid" md:"-"`
	CompanyShortName string   `json:"companyShortName" md:"股票名称"`
	CompanyLogo      string   `json:"companyLogo,omitempty" md:"-"`
	BoardType        []string `json:"boardType" md:"-"`
	PubDate          string   `json:"pubDate" md:"发布时间"`
	UpdateDate       string   `json:"updateDate" md:"-"`
	Author           string   `json:"author" md:"-"`
	AuthorName       string   `json:"authorName" md:"-"`
	PubClient        string   `json:"pubClient" md:"-"`
	AttachedId       string   `json:"attachedId" md:"-"`
	AttachedContent  string   `json:"attachedContent" md:"上市公司回复"`
	AttachedAuthor   string   `json:"attachedAuthor" md:"-"`
	AttachedPubDate  string   `json:"attachedPubDate" md:"回复时间"`
	Score            float64  `json:"score" md:"-"`
	TopStatus        int      `json:"topStatus" md:"-"`
	PraiseCount      int      `json:"praiseCount" md:"-"`
	PraiseStatus     bool     `json:"praiseStatus" md:"-"`
	FavoriteStatus   bool     `json:"favoriteStatus" md:"-"`
	AttentionCompany bool     `json:"attentionCompany" md:"-"`
	IsCheck          string   `json:"isCheck" md:"-"`
	QaStatus         int      `json:"qaStatus" md:"-"`
	PackageDate      string   `json:"packageDate" md:"-"`
	RemindStatus     bool     `json:"remindStatus" md:"-"`
	InterviewLive    bool     `json:"interviewLive" md:"-"`
}

type CailianpressWeb struct {
	Total int `json:"total"`
	List  []struct {
		Title   string `json:"title" md:"资讯标题"`
		Ctime   int    `json:"ctime" md:"资讯时间"`
		Content string `json:"content" md:"资讯内容"`
		Author  string `json:"author" md:"资讯发布者"`
	} `json:"list"`
}

type BKDict struct {
	gorm.Model  `md:"-"`
	BkCode      string `json:"bkCode" md:"行业/板块代码"`
	BkName      string `json:"bkName" md:"行业/板块名称"`
	FirstLetter string `json:"firstLetter" md:"first_letter"`
	FubkCode    string `json:"fubkCode" md:"fubk_code"`
	PublishCode string `json:"publishCode" md:"publish_code"`
}

func (b BKDict) TableName() string {
	return "bk_dict"
}

type WordAnalyze struct {
	gorm.Model
	DataTime *time.Time `json:"dataTime" gorm:"index;autoCreateTime"`
	WordFreqWithWeight
}

// WordFreqWithWeight 词频统计结果，包含权重信息
type WordFreqWithWeight struct {
	Word      string
	Frequency int
	Weight    float64
	Score     float64
}

// SentimentResult 情感分析结果类型
type SentimentResult struct {
	Score         float64       // 情感得分
	Category      SentimentType // 情感类别
	PositiveCount int           // 正面词数量
	NegativeCount int           // 负面词数量
	Description   string        // 情感描述
}

type SentimentResultAnalyze struct {
	gorm.Model
	DataTime *time.Time `json:"dataTime" gorm:"index;autoCreateTime"`
	SentimentResult
}

// SentimentType 情感类型枚举
type SentimentType int

type HotStrategy struct {
	ChgEffect bool               `json:"chgEffect"`
	Code      int                `json:"code"`
	Data      []*HotStrategyData `json:"data"`
	Message   string             `json:"message"`
}

type HotStrategyData struct {
	Chg       float64 `json:"chg" md:"平均涨幅(%)"`
	Code      string  `json:"code" md:"-"`
	HeatValue int     `json:"heatValue" md:"热度值"`
	Market    string  `json:"market" md:"-"`
	Question  string  `json:"question" md:"选股策略"`
	Rank      int     `json:"rank" md:"-"`
}

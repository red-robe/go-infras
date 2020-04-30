package starter

// 应用程序启动管理器
type BootApplication struct {
	Sctx *SystemContext
}

// 启动选项
type StarterOption struct {
	StarterKey string      //该启动选项适配的启动器
	Resource   interface{} //该启动器适配的资源
}

func NewStarterOption(key string, res interface{}) *StarterOption {
	sopt := new(StarterOption)
	sopt.StarterKey = key
	sopt.Resource = res
	return sopt
}

// 创建应用程序启动管理器
func NewApplication(opts ...*StarterOption) *BootApplication {
	//创建启动管理器
	b := new(BootApplication)
	b.Sctx = &SystemContext{}
	//为启动器设置资源选项
	b.Sctx.SetStarterOptions(opts...)
	return b
}

// 启动应用程序所有基础资源 （初始化 -> 安装 -> 启动）
func (b *BootApplication) Start() {
	b.init()
	b.setup()
	b.start()
}

// 停止或销毁应用程序所有基础资源
func (b *BootApplication) Stop() {
	for _, s := range StarterManager.GetAll() {
		s.Stop()
	}
}

func (b *BootApplication) init() {
	for _, s := range StarterManager.GetAll() {
		s.Init(b.Sctx)
	}
}

func (b *BootApplication) setup() {
	for _, s := range StarterManager.GetAll() {
		s.Setup(b.Sctx)
	}
}

func (b *BootApplication) start() {
	for i, s := range StarterManager.GetAll() {

		if s.SetStartBlocking() { // 阻塞启动
			// 如果是最后一个可阻塞的，直接启动并阻塞
			if i == StarterManager.Len()-1 {
				s.Start(b.Sctx)
			} else {
				// 如果不是，使用goroutine来异步启动，防止阻塞后面starter
				go s.Start(b.Sctx)
			}
		} else { // 非阻塞启动
			s.Start(b.Sctx)
		}
	}
}

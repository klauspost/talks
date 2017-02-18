	// Provide a testing context // HL
	func NewTestContext(t *testing.T) context.Context

	// Return a logger that will print context values by default. // HL
	func Logger(c context.Context) loggers.Advanced 
	func WithField(parent context.Context, key string, val interface{}) context.Context 

	// Send an error to our error handling system. // HL
	func NoticeError(c context.Context, cause error, msg string) 

	// Provide benchmark functions that logs benchmarks with context. // HL
	func BenchSegment(ctx context.Context, name string) (stop func()) 
	func BenchHTTP(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) 
	func BenchRoundTrip(ctx context.Context, client *http.Client) *http.Client 
	func BenchGormDB(c context.Context, db *gorm.DB) *gorm.DB 

	// Ignores parent cancellation. // HL
	func AsyncBg(parent context.Context) context.Context 

package awsctx

type AwsRequest struct {
	Service string
	Action  string
	Input   interface{}
	Output  interface{} // this is not populated until after the api call has executed
	Error   error       // either output _or_ error are populated
}

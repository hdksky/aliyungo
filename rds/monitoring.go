package rds

import (
	"github.com/hdksky/aliyungo/common"
	"github.com/hdksky/aliyungo/util"
)



type DescribeDBInstancePerformanceArgs struct{
	DBInstanceId	string
	key		string
	StartTime	string
	EndTime		string
}


type PerformanceValueType  struct{
	Value       	string
	Date        	util.ISO6801Time
}

type PerformanceKeyType struct {
	Key         	string
	Unit        	string
	ValueFormat 	string
	Values      struct {
			    PerformanceValue []PerformanceValueType
		    }
}

type DescribeDBInstancePerformanceResponse struct {
	common.Response
	DBInstanceId 	string
	Engine 		string
	StartTime 	util.ISO6801Time
	EndTime 	util.ISO6801Time
	PerformanceKeys struct {
			     PerformanceKey []PerformanceKeyType
		     }
}


func (client *DescribeDBInstancePerformanceArgs) Setkey(key string)  {
	client.key = key
}


func (client *Client) DescribeDBInstancePerformance(args *DescribeDBInstancePerformanceArgs) (resp DescribeDBInstancePerformanceResponse, err error) {

	response := DescribeDBInstancePerformanceResponse{}
	err = client.Invoke("DescribeDBInstancePerformance", args, &response)
	if err != nil {
		return response, err
	}
//	fmt.Print("zheshigesha  ",response,"\n")
	return response, nil

}

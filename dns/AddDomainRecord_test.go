package dns

import (
	"testing"
)

func TestAddDomainRecord(t *testing.T) {
	client := NewTestClient()
	rr := "*.ce4e2104fd10e4cd68ce21c40c28805e8.cn-beijing"

	//describe
	records,err := describeDomainRecords(t,client,rr)
	if err!=nil{
		t.Fatalf("Error %++v",err )
	}else{
		t.Logf("Records: %#v",records)
	}

	//add
	//ip := "47.95.138.155"
	//addDomainRecord(t,client,rr,ip)

	//deleted
	//for _,record := range records.DomainRecords.Record{
	//	deleteDomainRecord(t,client,record.RecordId)
	//}
}

func describeDomainRecords(t *testing.T,client *Client, rr string)(*DescribeDomainRecordsResponse,error) {
	describeDomainRecordsArgs := DescribeDomainRecordsArgs{
		DomainName:TestDomainName,
		RRKeyWord:   rr,
		TypeKeyWord: ARecord,
	}

	response, err := client.DescribeDomainRecords(&describeDomainRecordsArgs)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		for _, item := range response.DomainRecords.Record {
			t.Logf("Record : %#v", item)
		}
	}

	return response,err
}

func addDomainRecord(t *testing.T,client *Client, rr, ip string)(*AddDomainRecordResponse,error ) {
	addDomainRecordArgs := AddDomainRecordArgs{
		DomainName: TestDomainName,
		RR:         rr,
		Type:       ARecord,
		Value:      ip,
	}
	response, err := client.AddDomainRecord(&addDomainRecordArgs)
	if err!=nil{
		t.Errorf("Error %#v",err )
		return nil,err
	}else{
		t.Logf("Response %#v",response)
	}

	return response,nil
}


func deleteDomainRecord(t *testing.T,client *Client,recordId string)(*DeleteDomainRecordResponse,error ){
	deleteDomainRecordArgs := DeleteDomainRecordArgs{
		RecordId:recordId,
	}

	response,err := client.DeleteDomainRecord(&deleteDomainRecordArgs)
	if err!=nil{
		t.Errorf("Error %#v",err)
	}else{
		t.Logf("Success")
	}

	return response,err
}
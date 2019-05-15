pragma solidity ^0.4.17;
//pragma experimental ABIEncoderV2;
import "lib/strings.sol";
//import "github.com/Arachnid/solidity-stringutils/strings.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract InHospitalData is Ownable {
  using strings for *;
  struct InHospitalInfo {
    string timeStamp;
    string hospitalId;
    string devId;
    string personId;
    string personVeinData;
    string personPicData;
    string additionalInfo;
  }
  mapping(string => InHospitalInfo[]) data;

  event dataPushed(string _personid, string _timestamp, string _hospitalid, string _devid);

  constructor() public {
  }

  function getDataByPersonId(string _personid) public view returns(string) {
    string memory res;
    string memory doubleQuotationMark;  // "
    string memory colonMark;  // :
    string memory dcolondMark;    // ":"
    string memory dcommaMark;     // ","
    doubleQuotationMark = '"';
    colonMark = ':';
    dcolondMark = '":"';
    dcommaMark = '","';

    res = "[";

    for (uint i = 0; i < data[_personid].length; i++) {
      string memory itemData;

      itemData = "{";
      itemData = itemData.toSlice().concat(doubleQuotationMark.toSlice());

      itemData = itemData.toSlice().concat("timeStamp".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].timeStamp.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("hospitalId".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].hospitalId.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("devId".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].devId.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("personId".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].personId.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("personVeinData".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].personVeinData.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("personPicData".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].personPicData.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("additionalInfo".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].additionalInfo.toSlice());

      itemData = itemData.toSlice().concat(doubleQuotationMark.toSlice());
      itemData = itemData.toSlice().concat("}".toSlice());

      res = res.toSlice().concat(itemData.toSlice());
      if (i < data[_personid].length-1) {
        res = res.toSlice().concat(",".toSlice());
      }
    }

    res = res.toSlice().concat("]".toSlice());
    return res;
  }

  function getDataByPersonidAndDate(string _personid, string _date) public view returns(string) {
    string memory res;
    string memory doubleQuotationMark;  // "
    string memory colonMark;  // :
    string memory dcolondMark;    // ":"
    string memory dcommaMark;     // ","
    doubleQuotationMark = '"';
    colonMark = ':';
    dcolondMark = '":"';
    dcommaMark = '","';

    res = "[";

    uint flag = 0;
    for (uint i = 0; i < data[_personid].length; i++) {
      if (!data[_personid][i].timeStamp.toSlice().startsWith(_date.toSlice())) continue;

      string memory itemData;

      itemData = "{";
      itemData = itemData.toSlice().concat(doubleQuotationMark.toSlice());

      itemData = itemData.toSlice().concat("timeStamp".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].timeStamp.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("hospitalId".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].hospitalId.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("devId".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].devId.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("personId".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].personId.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("personVeinData".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].personVeinData.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("personPicData".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].personPicData.toSlice());
      itemData = itemData.toSlice().concat(dcommaMark.toSlice());
      itemData = itemData.toSlice().concat("additionalInfo".toSlice());
      itemData = itemData.toSlice().concat(dcolondMark.toSlice());
      itemData = itemData.toSlice().concat(data[_personid][i].additionalInfo.toSlice());

      itemData = itemData.toSlice().concat(doubleQuotationMark.toSlice());
      itemData = itemData.toSlice().concat("}".toSlice());

      if (flag>0) {
        itemData = ",".toSlice().concat(itemData.toSlice());
      }
      res = res.toSlice().concat(itemData.toSlice());
      flag = flag+1;
    }

    res = res.toSlice().concat("]".toSlice());
    return res;
  }

  function putData(string _key, string _timestamp, string _hospitalid, string _devid, string _personid, string _personveindata, string _personpicdata, string _additionalinfo) public onlyOwner {
    for (uint i = 0; i < data[_key].length; i++) {
      if (data[_key][i].timeStamp.toSlice().equals( _timestamp.toSlice())) return;
    }
    data[_key].push(
      InHospitalInfo({
        timeStamp:_timestamp,
        hospitalId:_hospitalid,
        devId:_devid,
        personId:_personid,
        personVeinData:_personveindata,
        personPicData:_personpicdata,
        additionalInfo:_additionalinfo})
    );

    emit dataPushed(_personid, _timestamp, _hospitalid, _devid);
  }
}

class koke_class {
    constructor(url,label) {
      this.ss = SpreadsheetApp.getActiveSpreadsheet();
      this.sheet = this.ss.getActiveSheet();
      this.url = url;
      this.label = label;
      this.data = UrlFetchApp.fetch(this.url).getContentText();
      this.data = JSON.parse(this.data);
      this.data = this.data.data.sort();
    }
  
    SelectSheet(name) { 
      let sheet = this.ss.getSheetByName(name);
      return sheet;
    }
    InsertSheet(name) { 
      let newSheet = this.ss.insertSheet();
      newSheet.setName(name);
    }
    
    InsertData(name) {
      InsertSheet(name);
      this.sheet = SelectSheet(name);
      
      this.label.forEach((l,index) => {
      //sheet.getRange(0,index).setValue(l);
      console.log(l,index);
      this.sheet.getRange(1,index+1).setValue(l);
      })
  
      for(let d of this.data) {
        let lastRow = this.sheet.getLastRow();
        console.log(d.ImageName);
        console.log(lastRow);
        this.sheet.getRange(lastRow+1,1).setValue(d.ImageName);
        this.sheet.getRange(lastRow+1,2).setValue(d.Lng);
        this.sheet.getRange(lastRow+1,3).setValue(d.Lat);
      }
    }
  };
  
  let class_run = (number) => {
    let url = `https://peaceful-temple-64943.herokuapp.com/koke/${number}`;
    let name = `擁壁${number}`;
    let label = ["写真名","緯度","経度","種名"];
  
    let koke = new koke_class(url,label);
    console.log(koke.url);
    //console.log(koke.data);
    koke.InsertData(name);
  }
  
  let  main = () => {
    class_run(20);
  }
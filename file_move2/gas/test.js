class koke_class {
  constructor(url,label) {
    this.ss = SpreadsheetApp.getActiveSpreadsheet();//ここで使うスプレッドシート
    this.sheet = null;//書き込むシート
    this.url = url;//サーバーのURL
    this.label = label;//ラベル
    this.data = null;//データ
    this.code = null;//サーバーステータスコードを格納（いらない気もするが
    this.err = null;//エラー格納。正常時はnullを想定
  }

  //データを用意
  SetData() {
    
    try {
      
      let res = UrlFetchApp.fetch(this.url);
      this.code = res.getResponseCode();
      //サーバー側からのステータスコード
      if(this.code == 201 ){ 
        let data = res.getContentText();
        data = JSON.parse(data);
        this.data = data.data.sort();
      }

    }catch(e){
      Logger.log('Error:')
      Logger.log(e)
      this.err = e;

    } finally{ 
      return this.code;
    }

  }

  //どのシートを選ぶかを決める
  SelectSheet(name) { 
    let sheet = this.ss.getSheetByName(name);
    return sheet;
  }
  
  //新しいシートを挿入
  InsertSheet(name) { 
    let newSheet = this.ss.insertSheet();
    newSheet.setName(name);
  }
  
  //レスからのデータを入力
  InsertData(name) {
    SetData();
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

//typedef super koke_class

class kokeClass extends koke_class {
constructor(url,label,name) {
  
  super(url,label);
  this.name = name;
}

//override 
//エラーの時は、何もせずに終わる
InsertData() { 

  super.SetData();

  if(this.err == null) {
    
    console.log(this.name);
    super.SetData();
    super.InsertSheet(this.name);
    this.sheet = super.SelectSheet(this.name);
    
    this.label.forEach((l,index) => {
      this.sheet.getRange(1,index+1).setValue(l);
    })

    for(let d of this.data) {
      let lastRow = this.sheet.getLastRow();
      this.sheet.getRange(lastRow+1,1).setValue(d.ImageName);
      this.sheet.getRange(lastRow+1,2).setValue(d.Lng);
      this.sheet.getRange(lastRow+1,3).setValue(d.Lat);
    }

  }else {
    console.log(`Error Message ${this.err}`);
  }
}

}

//実行するためのクラス
//gasは、関数単位で実行できるが、複数あると気づかずに別の関数を実行してしまうことがある。
//そのため静的メソッドにより実行している。
class Run {
static run(number) {
  //https://localhost:3000/koke/
  let url = `https://localhost:3000/koke/${number}`;
  let name = `擁壁${number}`;
  let label = ["写真名","緯度","経度","種名"];

  console.log(name);
  let koke = new kokeClass(url,label,name);
  console.log(koke.url);
  koke.InsertData();
}
}

let  main = () => {
  for(let i = 1; i<=20; i++) {
    //なぜか１０だけない
    Run.run(i);
  }
}
#include <vector>
#include <string>
#include <fstream>
#include <iostream>
using ld = long double;

std::vector<ld> operator+(std::vector<ld>& x, std::vector<ld> &y) {
    std::vector<ld> res(x.size());
    for(int i = 0; i < (int)res.size(); ++i) res[i] = x[i]+y[i];
    return res;
}
std::vector<ld> operator/(std::vector<ld>& x, int y) {
    std::vector<ld> res(x.size());
    for(int i = 0; i < (int)res.size(); ++i) res[i]=x[i]/y;
    return res;
}

//文字列のsplit機能
std::vector<std::string> split(std::string str, char del) {
    int first = 0;
    //文字列中にdelが最初に出るところを返す，
    int last = str.find_first_of(del);
    std::vector<std::string> result;
    while (first < str.size()) {
        //strのfirst から　last-first 
        std::string subStr(str, first, last - first);
        result.push_back(subStr);
        first = last + 1;
        //first 以降のdelが含まれている場所を最初に出るところを返す．
        last = str.find_first_of(del, first);
        //std::string::npos 見つからんかった時の返り値
        if (last == std::string::npos) {
            last = str.size();
        }
    }
    return result;
}

std::vector<std::vector<std::string> >
csvtovector(std::string filename, int labels = 0){
    std::ifstream reading_file;
    reading_file.open(filename, std::ios::in);
    if(!reading_file){
        std::vector<std::vector<std::string> > data;
        return data;
    }
    std::string reading_line_buffer;
    //最初のignore_line_num行を空読みする
    for(int line = 0; line < labels; line++){
        getline(reading_file, reading_line_buffer);
        if(reading_file.eof()) break;
    }

    std::vector<std::vector<std::string> > data;
    while(std::getline(reading_file, reading_line_buffer)){
        if(reading_line_buffer.size() == 0) break;
        std::vector<std::string> temp_data;
        temp_data = split(reading_line_buffer, ',');
        data.push_back(temp_data);
    }
    return data;
}
//平均計算
std::vector<std::vector<ld> > clac(std::string path,int interval) {
    std::vector<std::vector<std::string> > data = csvtovector(path,1);
    std::vector<std::vector<ld> > lddata(data.size());
    std::vector<std::vector<ld> > res(data.size(),std::vector<ld>(data[0].size(),0));
    for(int i = 0; i < data.size(); ++i) {
        for(int j = 0; j < data[i].size(); ++j) {
            lddata[i].push_back(std::stold(data[i][j]));
        }
    }
    for(int i = 0; i < data.size(); ++i) {
        for(int j = 0; j < data[i].size(); ++j) {
            std::cout << lddata[i][j] << " ";
        }
        std::cout << std::endl;
    }
    //演算子のオーバーロードを使用
    for(int i = 0; i < lddata.size(); ++i) {
        if(lddata.size() < i+interval) break;
        for(int j = i; j < i+interval; ++j) {
            res[i] = res[i] + lddata[j];
        }
        res[i]= res[i]/interval;
        i+=interval-1;
    }
    return res;
}

int main(){
    std::string filepath , create_filename;
    int clac_interval; 
    std::cout << "計算の対象のファイルのパスを入力してください．"<< std::endl;
    std::cin >> filepath;
    std::cout << "平均計算の間隔を入力してください" << std::endl;
    std::cin >> clac_interval;
    std::vector<std::vector<ld> > data = clac(filepath,clac_interval);
    std::cout<< "ファイル名を入力してください．(example.csv)" << std::endl;
    std::cin >> create_filename;
    std::ofstream ofs(create_filename);
    for(int i = 0; i < data.size(); ++i) {
        if(data[i][0] == 0) continue;
        for(int j = 0; j < data[i].size(); ++j) {
            ofs << data[i][j] << ",";
        }
        ofs << std::endl;
    }
    return 0;
}

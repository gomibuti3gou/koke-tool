import csv
import numpy as np
import pandas as pd
%matplotlib inline
import matplotlib.pyplot as plt
import japanize_matplotlib
 
 
def rowMeans(start,end,data,datacolumns):
  r = data[datacolumns[start]]
  for i in range(start+1,end+1):
    r += data[datacolumns[i]]
  return r/(end-start+1)

def amin(start,end,data):
  res = data[start:end]
  flag = 100000
  ans = 0
  #res = 0.0
  for i,value in enumerate(res):
    if flag > value:
      flag = value
      ans = i
  if(min(res) != flag): print("warn")
  #else : print("warn!!")
  print(ans+(start))
  return (start+ans,ans)

def amax(start,end,data):
  res = data[start:end]
  flag = -1
  ans = 0
  for i,value in enumerate(res):
    if flag < value:
      flag = value
      ans = i
  return start+ans+325
def Standart(list):
  return (list - min(list))/(max(list) - min(list))

def SCsv(data,number):
  return (data[number] -min(data[number]))/(max(data[number]) - min(data[number]))

def Standardization(number,data,df):
  return (data[number] - df.loc['min',number])/(df.loc['max',number] - df.loc['min',number])

def dfx(list):
  res = []
  for i in range(0,len(list)-1):
    if i == 0: res.append(0)
    else : res.append(list[i]-list[i-1])
  return res;
def dfx_standart(data):
  res = []
  for d in data:
    res.append((d-min(data)) /(max(data)-min(data)))
  return res;

df = pd.read_csv("./experiment-9-5.csv",header=None)
aa = df.values[:,1:]
headers = df.loc[0:,:0].astype('str')
handheld = pd.DataFrame(aa.T,columns = headers[0])
hand_columns = list(handheld.columns)
x = handheld[hand_columns[0]]
handheld.head()

y1 = rowMeans(94,103,handheld,hand_columns)
y1 = Standart(y1)
y2 = rowMeans(104,113,handheld,hand_columns)
y2 = Standart(y2)
y3 = rowMeans(114,123,handheld,hand_columns)
y3 = Standart(y3)
y4 = rowMeans(124,133,handheld,hand_columns)
y4 = Standart(y4)

ig, ax = plt.subplots()
ax.plot(x,y1,label="30度 (HandHeld2)")
ax.plot(x,y2,label="45度 (HandHeld2)")
ax.plot(x,y3,label="60度 (HandHeld2)")
ax.plot(x,y4,label="90度 (HandHeld2)")
ax.legend()
plt.title("角度による波長強度の差異(正規化)")
plt.xlabel("波長 λ(nm)")
plt.ylabel("反射強度 (エゾノギシギシ)")
plt.grid()
plt.savefig("./image/kakudo1.png")
plt.show()

y1 = rowMeans(94,103,handheld,hand_columns)
y2 = rowMeans(104,113,handheld,hand_columns)
y3 = rowMeans(114,123,handheld,hand_columns)
y4 = rowMeans(124,133,handheld,hand_columns)

ig, ax = plt.subplots()
ax.plot(x,y1,label="30度 (HandHeld2)")
ax.plot(x,y2,label="45度 (HandHeld2)")
ax.plot(x,y3,label="60度 (HandHeld2)")
ax.plot(x,y4,label="90度 (HandHeld2)")
ax.legend()
plt.title("角度による波長強度の差異")
plt.xlabel("波長 λ(nm)")
plt.ylabel("反射強度 (エゾノギシギシ)")
plt.grid()
plt.savefig("./image/kakudo2.png")
plt.show()

y1 = rowMeans(134,143,handheld,hand_columns)
y1 = Standart(y1)
y2 = rowMeans(144,153,handheld,hand_columns)
y2 = Standart(y2)
y3 = rowMeans(154,163,handheld,hand_columns)
y3 = Standart(y3)
y4 = rowMeans(164,173,handheld,hand_columns)
y4 = Standart(y4)

ig, ax = plt.subplots()
ax.plot(x,y1,label="10cm (HandHeld2)")
ax.plot(x,y2,label="20cm (HandHeld2)")
ax.plot(x,y3,label="30cm (HandHeld2)")
ax.plot(x,y4,label="40cm (HandHeld2)")
ax.legend()
plt.title("距離による波長強度の違い(正規化）")
plt.xlabel("波長 λ(nm)")
plt.ylabel("反射強度 (エゾノギシギシ)")
plt.grid()
plt.savefig("./image/kyori1.png")
plt.show()

y1 = rowMeans(134,143,handheld,hand_columns)
y2 = rowMeans(144,153,handheld,hand_columns)
y3 = rowMeans(154,163,handheld,hand_columns)
y4 = rowMeans(164,173,handheld,hand_columns)

ig, ax = plt.subplots()
ax.plot(x,y1,label="10cm (HandHeld2)")
ax.plot(x,y2,label="20cm (HandHeld2)")
ax.plot(x,y3,label="30cm (HandHeld2)")
ax.plot(x,y4,label="40cm (HandHeld2)")
ax.legend()
plt.title("距離による波長強度の違い")
plt.xlabel("波長 λ(nm)")
plt.ylabel("反射強度 (エゾノギシギシ)")
plt.grid()
plt.savefig("./image/kyori2.png")
plt.show()
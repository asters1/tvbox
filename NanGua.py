import requests
import hashlib
import json
from bs4 import BeautifulSoup
from urllib.parse import urlparse
from urllib.parse import quote
import sys
import sys
import time

extend=""
appVersionCode = "9"
appVersionName = "1.0.9"
deviceModel="M2012K11AC"
deviceVersion="14"
deviceBrand="Redmi"



def init(etd):
    global extend
    extend=etd
def homeContent(filter):
    TX={
            "type_id":"TX",
            "type_name":"TX"
            }
    IQY={
            "type_id":"IQY",
            "type_name":"IQY"

            }
    MG={
            "type_id":"MG",
            "type_name":"MG"

            }
    classes=[TX,IQY,MG]
    filter={"TX":GetTXFiltter()}
    result={
            "class":classes,
            "filter":filter

            }
    jstr=json.dumps(result,ensure_ascii=False)
    print(jstr)
def categoryContent(tid,pg,filter,extend):
    result={}
    if(tid=="TX"):
        SX=""
        json_filter=json.loads(extend)
        for k,v in json_filter.items():
            # print(f)
            SX=SX+"&"+k+"="+v
        url="https://v.qq.com/x/bu/pagesheet/list?_all=1&append=1&channel=child&listpage=1&offset="+str(((int(pg) - 1) * 21))+"&pagesize=21&sort=75"+SX;
        res_content=requests.get(url,GetHeaders()).text.encode("iso-8859-1").decode("utf-8")
        sp=BeautifulSoup(res_content,features="html.parser")
        list_items=sp.find_all(class_="list_item")
        list=[]
        for item in list_items:
            vod_remarks_item=item.find_all(class_="figure_caption")
            vod_remarks=""
            vod_pic=Completion(url,item.a.img["src"])
            if len(vod_remarks_item)>0:
                vod_remarks=vod_remarks_item[0].text

            vod={
                    "vod_id" : item.a["title"],
                    "vod_name" : item.a["title"],
                    "vod_remarks":vod_remarks,
                    "vod_pic":vod_pic
            }
            list.append(vod)
            result={
                    "page":pg,
                    "pagecount":sys.maxsize,
                    "limit":90,
                    "total":sys.maxsize,
                    "list":list
                    }
        
    jstr=json.dumps(result,ensure_ascii=False)
    print(jstr)

        
def detailContent(ids):
    ti=GetTime()
    url="http://ys.changmengyun.com/api.php/provide/vod_detail?appVersionName="+appVersionName+"&imei=&time="+ti+"&id="+ids+"&deviceScreen=2340*1080&appVersionCode="+appVersionCode+"&deviceModel"+deviceModel+"&app=ylys&deviceBrand="+deviceBrand+"&devices=android&deviceVersion="+deviceVersion
    res=requests.get(url,headers=GetNGHeaders(ti)).text
    json_res=json.loads(res)
    data_obj=json_res["data"]
    player_info_array=data_obj["player_info"]
    # print(player_info_array)
    play_url_list=[]
    play_from_array=[]
    for v_obj in player_info_array:
        play_from_array.append(v_obj["show"])
        u_array=v_obj["video_info"]
        url_list=[]
        for u in u_array:
            url_list.append(u["name"]+"$"+u["url"][0])
        # print(url_list)
        play_url_list.append("#".join(url_list))
        # print(play_url_list)

    vod_obj={
        "vod_id":ids,
        "vod_name":data_obj["name"],
        "vod_pic":data_obj["img"],
        "type_name":data_obj["type"],
        "vod_year":data_obj["year"],
        "vod_remarks":data_obj["msg"],
        "vod_content":data_obj["info"],
        "vod_play_from":"$$$".join(play_from_array),
        "vod_play_url":"$$$".join(play_url_list)
            }
    result={
            "list":[vod_obj]
            }
    jstr=json.dumps(result,ensure_ascii=False)
    print(jstr)

def searchContent(key):
    list=[]
    ti=GetTime()
    url="http://ys.changmengyun.com/api.php/provide/search_result?video_name="+quote(key)+"&appVersionName="+appVersionName+"&imei=&time="+ti+"&deviceScreen=2340*1080&appVersionCode="+appVersionCode+"&deviceModel="+deviceModel+"&devices=android&deviceVersion="+deviceVersion
    res=requests.get(url,headers=GetNGHeaders(ti)).text
    json_res=json.loads(res)
    arrayres=json_res["data"][0]["data"]
    for item in arrayres:
        vod={
            "vod_id":str(item["id"]),
            "vod_name":item["video_name"],
            "vod_pic":item["img"],
            "vod_remarks":item["category"]
                }
        list.append(vod)
    result={"list":list}
    # print(result)
    jstr=json.dumps(result,ensure_ascii=False)
    print(jstr)
def playerContent(flag ,id):
    ti=GetTime()
    
    res_content=requests.get(id,headers=GetNGHeaders(ti)).text
    json_res=json.loads(res_content)
    result={
        "parse":0,
        "header":json_res["data"]["header"],
        "url":json_res["data"]["url"]

            }
    jstr=json.dumps(result,ensure_ascii=False)
    print(jstr)


    # homeContent(True)




#===============自定义函数
def GetNGHeaders(ti):
    headers={
        "user-agent":"okhttp/3.12.0",
        "version_name":appVersionName,
        "version_code":appVersionCode,
        "sign":GetMd5("#uBFszdEM0oL0JRn@"+ti),
        "timeMillis":ti
            }
    return headers

def GetMd5(text):
    md5=hashlib.md5()
    md5.update(text.encode("utf-8"))
    return md5.hexdigest()
def GetTime():
    t=str(int(time.time()*1000))
    return t

def Completion(str1,str2):
    result=""
    parse=urlparse(str1)
    if str2.startswith("//"):
        result=parse.scheme+":"+str2
    elif str2.startswith("://"):
        result=parse.scheme+str2
    elif str2.startswith("http://") or str2.startswith("https://"):
        result=str2
    else:
        result=parse.scheme+"://"+str2
    return result
def GetHeaders():
    headers={
        "user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
    }
    return headers
def GetTXFiltter():
    #地区
    iarea_value=[{"n":"全部","v":"-1"},{"n":"国内","v":"3"},{"n":"欧美","v":"1"},{"n":"日韩","v":"2"}]
    iarea={"key":"iarea","name":"地区","value":iarea_value}
    #年龄
    iyear_value=[{"n":"全部","v":"-1"},{"n":"0-3岁","v":"1"},{"n":"4-6岁","v":"2"},{"n":"7-9岁","v":"3"},{"n":"10岁以上","v":"4"}]
    iyear={"key":"iyear","name":"年龄","value":iyear_value}
    #性别
    gender_value=[{"n":"全部","v":"-1"},{"n":"男孩","v":"2"},{"n":"女孩","v":"1"},]
    gender={"key":"gender","name":"性别","value":gender_value}
    #类型
    itype_value=[{"n":"全部","v":"-1"},{"n":"玩具","v":"4"},{"n":"交通工具","v":"10"},{"n":"手工·绘画","v":"3"},{"n":"儿歌","v":"1"},{"n":"益智早教","v":"2"},{"n":"英语","v":"5"},{"n":"早教","v":"6"},{"n":"数学","v":"7"},{"n":"国学","v":"8"},{"n":"冒险","v":"9"},{"n":"魔幻·科幻","v":"11"},{"n":"动物","v":"12"},{"n":"真人·特摄","v":"13"},{"n":"探索","v":"14"},{"n":"其他","v":"15"}]
    itype={"key":"itype","name":"类型","value":itype_value}
    TX=[iarea,iyear,gender,itype]

    # jstr=json.dumps(TX,ensure_ascii=False)
    return TX


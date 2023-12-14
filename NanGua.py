import requests
import json
from bs4 import BeautifulSoup
from urllib.parse import urlparse
import sys
import sys

extend=""
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
    print("{}")
def searchContent(key):
    print("{}")
def playerContent(flag ,id):
    print("{}")
    # homeContent(True)




#===============自定义函数
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
    # print(jstr)
# GetTXFiltter()
# homeContent(True)
# categoryContent("TX",2,True,'{"iyear":"1"}')

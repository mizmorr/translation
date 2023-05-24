file=open("text.txt","r")
rpn=file.read()
def replace_ids(num,filepath):
    ids=open("data/task"+str(num)+"/idents.txt","r").readlines()
    count=0
    file=open(filepath,"r").read()
    for element in file:
        count+=element.count("I")
    print(count)
    res=""
    if count!=0:
        res=file.replace("I1",ids[0])
    for i in range (count-1):
        if res.split("\n")[res.split("\n").index("I"+str(i+2))-1].startswith("BC"):
            res=res.replace("I"+str(i+2),"class "+ids[1]+":")

        elif res.split("\n")[res.split("\n").index("I"+str(i+2))-1].startswith("BF"):
            res=res.replace("I"+str(i+2),"def "+ids[1]+"(")
        else: res=res.replace("I"+str(i+2),ids[1])

def replace_c(num,filepath):
    ids=open("data/task"+str(num)+"/c.txt","r").readlines()
    count=0
    file=open(filepath,"r").read()
    for element in file:
        count+=element.count("C")
    print(count)
    res=""
    if count!=0:
        res=file.replace("C1",ids[0])
    for i in range (count-1):
        res=res.replace("C"+str(i+2),ids[1])
    print(res)

def replace_num(num,filepath):
    ids=open("data/task"+str(num)+"/n.txt","r").readlines()
    count=0
    file=open(filepath,"r").read()
    for element in file:
        count+=element.count("N")
    print(count)
    res=""
    if count!=0:
        res=file.replace("N1",ids[0])
    for i in range (count-1):
        res=res.replace("N"+str(i+2),ids[1])
def keywords_replace(num,filepath):
    ids=open("data/task"+str(num)+"/n.txt","r").readlines()
    count=0
    file=open(filepath,"r").read()
    for element in file:
        count+=element.count("W")
    print(count)
    res=""
    if count!=0:
        res=file.replace("W1",ids[0])
    for i in range (count-1):
        res=res.replace("W"+str(i+2),ids[1])

def process(filepath,pathdest):
    rpn=open(filepath,"r").read()
    result=rpn.replace("BG","").replace("2AEA","").replace("STR","").replace("INT","").replace("LONG","").replace("END","").replace("]","").replace("[","")
    result=replace_ids(result)
    result=replace_num(result)
    result=replace_c(result)
    file=open(pathdest,"w")
    file.write(result)
    file.close()

n=2

for k in range(2,45):
    print(k)

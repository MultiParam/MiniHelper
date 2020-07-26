'''
Func: Count the words
Author: DawnGuo
'''

def addWord(word,wordCountDict):
    '''Update the word frequency:word is the key,frequency is the value.'''

    if word in wordCountDict:
        wordCountDict[word] += 1 #如果单词在wordCountDict词典已存在，数量加1
    else:
        wordCountDict[word] = 1 #如果不存在，插入一条记录

def processLine(line,wcDict):
    '''process each line of text'''

    line = line.strip() #去掉每行的空白
    wordList = line.split()  #去掉空白字符，分割字符

    word = wordList[0].lower()
    word = word.strip()

    addWord(word,wcDict)


def prettyPrint(wcDict):
    '''print the result'''
    valKeyList = []

    for key,val in wcDict.items():
        valKeyList.append((val,key))

    valKeyList.sort(reverse=True)
    print('%-10s%10s' %('Word','Count'))
    print('_'*21)
    for val,key in valKeyList:
        print('%-12s     %3d' %(key,val))

def main():
    wcDict = {}
    fobj = open('words.txt','r', encoding='utf8')
    for line in fobj:
        processLine(line,wcDict)
    prettyPrint(wcDict)


if __name__ == "__main__":
    main()
import string
f = open('names.txt')
names = f.readlines()
f.close()
nameList = names[0].split(",")
score = 0
nameList.sort()
for i in range(len(nameList)):
    nameScore = 0
    for c in nameList[i]:
        if c != "\"":
            nameScore += string.uppercase.index(c) + 1

    score += nameScore * (i + 1)

print score

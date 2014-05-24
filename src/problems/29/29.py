distinct = []
for a in range(2,101):
    for b in range(2,101):
        sum = a**b
        if sum not in distinct:
            distinct.append(sum)

print len(distinct)

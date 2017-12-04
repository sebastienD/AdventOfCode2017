
def computeLargeur(b1, b2, diag):
    mid = diag/2
    return max(b1, mid+b2) - min(b1, mid+b2)

def first(value):
    diag = 1
    longueur = 0
    while diag * diag < value:
        first = diag*diag + 1
        diag += 2
        longueur+=1
    sudest = diag * diag
    sudouest = sudest - (diag-1)
    nordouest = sudouest - (diag-1)
    nordest = nordouest - (diag-1)
    borne = 0
    
    if value in range(sudouest, sudest+1):
        print('bas')
        borne = sudouest
    elif value in range(nordouest, sudouest+1):
        print('ouest')
        borne = nordouest
    elif value in range(nordest, nordouest+1):
        print('haut')
        borne = nordest
    else: #if value in range(first-1, nordest + 1):
        print('est')
        borne = first-1

    largeur = computeLargeur(value, borne, diag)
    print('Long {0}, Larg {1}'.format(longueur, largeur))
    distance = longueur + largeur
    print('Distance {0}'.format(distance))
    return distance

class line:
    def __init__(self, size):
        self.size = size
        self.content = []

    def append(self, v):
        self.content.append(v)
        return self
    
    def get(self, index):
        if index < 0 or index >= self.size:
            return 0
        return self.content[index]
    
    def isFull(self):
        return len(self.content) == self.size

    def string(self):
        return '{0}'.format(self.content)

class square:
    def __init__(self, size):
        self.size = size
        self.est = line(size)
        self.ouest = line(size)
        self.nord = line(size)
        self.sud = line(size)

    def string(self):
        return 'n:{0} e:{1} s:{2} o:{3}'.format(self.nord.string(), self.est.string(), self.sud.string(), self.ouest.string())

    def estNext(self, back, i):
        v = back.est.get(i-2) + back.est.get(i-1) + back.est.get(i) + self.est.get(i-1)
        #print('est {0}'.format(v))
        self.est.append(v)
        return v

    def nordNext(self, back, i):
        b = (back.nord.get(i-2) if (back.nord.get(i-2) != 0) else self.est.get(self.size-2))
        #print('nord {0} {1} {2}'.format(back.nord.get(i-2), self.est.get(self.size-2), b))
        v = b + back.nord.get(i-1) + back.nord.get(i) + self.nord.get(i-1)
        self.nord.append(v)
        return v

    def ouestNext(self, back, i):
        b = back.ouest.get(i-2) if (back.ouest.get(i-2) != 0) else self.nord.get(self.size-2)
        v = b + back.ouest.get(i-1) + back.ouest.get(i) + self.ouest.get(i-1)
        self.ouest.append(v)
        return v
    
    def sudNext(self, back, i):
        b = back.sud.get(i-2) if (back.sud.get(i-2) != 0) else self.ouest.get(self.size-2)
        f = back.sud.get(i) if (back.sud.get(i) != 0) else self.est.get(1)
        v = b + back.sud.get(i-1) + f + self.sud.get(i-1)
        #print('sud {0} {1} {2}'.format(back.sud.get(i-2), self.ouest.get(self.size-2), b))
        self.sud.append(v)
        return v

def second(waitFor):
    size = 1

    first = square(size)
    first.est.append(1)
    first.ouest.append(1)
    first.nord.append(1)
    first.sud.append(1)
    print(first.string())

    current = first
    last=0

    class Found(Exception): pass
    try:
        while last < waitFor:
            back = current
            size += 2
            current = square(size)

            current.est.append(0)
            index = 1
            while not current.est.isFull():
                last = current.estNext(back, index)
                if last > waitFor:
                    raise Found
                index += 1
            current.nord.append(last)

            index = 1
            while not current.nord.isFull():
                last = current.nordNext(back, index)
                if last > waitFor:
                    raise Found
                index += 1
            current.ouest.append(last)

            index = 1
            while not current.ouest.isFull():
                last = current.ouestNext(back, index)
                if last > waitFor:
                    raise Found
                index += 1
            current.sud.append(last)

            index = 1
            while not current.sud.isFull():
                last = current.sudNext(back, index)
                if last > waitFor:
                    raise Found
                index += 1
            current.est.content[0] = last

            print(current.string())
    except Found:
        print ('Found')
    return last

# print(first(277678))
print(second(277678))
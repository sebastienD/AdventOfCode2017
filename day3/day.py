
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
    #print 'Diag {0} {1} {2} {3}'.format(sudest, sudouest, nordouest, nordest)
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

    def isFull(self):
        return len(self.nord) == self.size and len(self.est) == self.size and len(self.sud) == self.size and len(self.ouest) == self.size

    def first(self, back):
        self.est.append(0)
        self.append(back.est[1])

# print(first(277678))
size = 1

s = square(size)
s.est.append(1)
s.ouest.append(1)
s.nord.append(1)
s.sud.append(1)
print(s.string())
print(s.isFull())

back=s
size+=1
current = square(size)
current.first(back)
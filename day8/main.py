from operator import gt,eq,ne,lt,le,ge,add,sub

def initIfNecessary(name1, name2):
    if name1 not in registers.keys():
        registers[name1]=0
    if name2 not in registers.keys():
        registers[name2]=0

def registerByName(name):
    if name in registers.keys():
        return registers[name]
    return 0

def analyzeLine(line):    
    elements = line.split(' ')
    initIfNecessary(elements[0], elements[4])
    cond = evalCondition(elements[4], elements[5], elements[6])
    print("cond {}".format(cond))
    if cond:
       return evalAction(elements[0], elements[1], elements[2])
    return 0

def evalCondition(nameRegister, ope, value):
    register = registerByName(nameRegister)
    print("operateur {} ".format(ope))
    if ope == ">":
        return gt(register, int(value))
    elif ope == "<":
        return lt(register, int(value))
    elif ope == "==":
        return eq(register, int(value))
    elif ope == "!=":
        return ne(register, int(value))
    elif ope == ">=":
        return ge(register, int(value))
    else:
        return le(register, int(value))
    return True

def evalAction(nameRegister, op, value):
    register = registerByName(nameRegister)
    print("regsiter {}".format(register))
    if eq(op,"inc"):
        register = add(register, int(value))
    else:
        register = sub(register, int(value))
    registers[nameRegister] = register
    print("regsiters {}".format(registers))
    return register

registers = {}
max4ever = 0

with open('input.txt') as lines:
    for line in lines:
        value = analyzeLine(line)
        if value > max4ever:
            max4ever = value

maxk = ''
maxv = 0
for r in registers:
    if registers[r] > maxv:
        maxk = r
        maxv = registers[r]

print(registers)
print(maxk, maxv, max4ever)
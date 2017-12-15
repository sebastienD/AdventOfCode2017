
fun round(listOfLength: List<Int>, numbers: MutableList<Int>, pair: Pair<Int, Int>): Pair<Int, Int> {
    val size = numbers.size

    var currentPosition = pair.first
    var skip = pair.second

    for (i in 0..(listOfLength.size-1)) {
        val indexMin = currentPosition
        var length = listOfLength[i]

        println("indexMin: "+indexMin+" length: "+length)

        val changedIndex = generateSequence(indexMin, { (it + 1) % size })
                .take(length)
                .toList()

        val changedValues = changedIndex
                .map { numbers[it] }
                .reversed()

        for (i in 0..(length - 1)) {
            numbers[changedIndex[i]] = changedValues[i]
        }

        currentPosition = (currentPosition + length + skip)%size
        skip++

        println("current position: %s, skip: %s".format(currentPosition, skip))
    }

    println(numbers)
    return Pair(currentPosition, skip)
}

fun hash(toEncoded: String): String {
    var numbers = generateSequence(0, { it + 1 }).take(256).toMutableList()

    var lengths = toEncoded
            .toCharArray()
            .map { it.toInt() }
            .toMutableList()
    lengths.addAll(listOf(17,31,73,47,23))

    var pair = Pair(0,0)
    for (i in 1..64) {
        pair = round(lengths, numbers, pair)
    }

    val sparseHash = numbers

    var finalListe = ArrayList<Int>(16)
    for (b in 0..15) {
        var v = sparseHash[0 + b*16]
        for (i in 1..15) {
            v = v.xor(sparseHash[i + b*16])
        }
        finalListe.add(v)
    }

    return finalListe.joinToString("") { Integer.toHexString(it).padStart(2, '0') }
}

fun main(args: Array<String>) {
    //val listOfLength = mutableListOf(34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167)
    //val liste = launch(listOfLength, 256)
    //println("le produit des 2 premières valeur est %s".format(liste[0] * liste[1]))

    val hashValue = hash("34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167")
    println(hashValue)
}
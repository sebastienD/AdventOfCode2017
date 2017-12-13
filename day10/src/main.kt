import java.util.ArrayList

fun launch(lengths: List<Int>, number: Int): List<Int> {
    val numbers = generateSequence(0, { it + 1 }).take(number).toMutableList()

    var currentPosition = 0
    var skipSize = 0

    for (i in 0..(number-1)) {
        val indexMin = currentPosition
        var length = lengths[i]

        println("indexMin: "+indexMin+" length: "+length)

        val seq: List<Int> = generateSequence(indexMin, { (it + 1) % number }).take(length).toList()
        println("seq: "+seq)
        val sub = seq.map { it -> numbers[it] }.reversed()
        println(sub)

        var loop = 0
        for (s in seq) {
            numbers[s] = sub[loop]
            loop++
        }

        currentPosition = (currentPosition + length + skipSize)%number
        skipSize++

        println("end")
        println(numbers)
    }

    println(numbers)

    return numbers
}

fun main(args: Array<String>) {

    val lengths: List<Int> = mutableListOf(34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167)

    launch(lengths, 256)
}
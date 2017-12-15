import kotlin.test.assertEquals
import org.junit.Test as test

class Main_testKt() {

    @test fun `should rerverse list`() {
        val lengths: List<Int> = mutableListOf(3, 4, 1, 5)
        val numbers = generateSequence(0, { it + 1 }).take(5).toMutableList()
        val pair = round(lengths, numbers, Pair(0,0))
        assertEquals(listOf(3,4,2,1,0) , numbers)
        assertEquals(4 , pair.first)
        assertEquals(4 , pair.second)
    }

    @test fun `number to xor`() {
        val v = 65.xor(27).xor(9).xor(1).xor(4)
                .xor(3).xor(40).xor(50).xor(91)
                .xor(7).xor(6).xor(0).xor(2)
                .xor(5).xor(68).xor(22)
        assertEquals(64 , v)
    }

    @test fun `empty hash`() {
        assertEquals("a2582a3a0e66e6e86e3812dcb672a272" , hash(""))
    }

    @test fun `AoC 2017 hash`() {
        assertEquals("33efeb34ea91902bb2f59c9920caa6cd" , hash("AoC 2017"))
    }

    @test fun `1,2,3 hash`() {
        assertEquals("3efbe78a8d82f29979031a4aa0b16a9d" , hash("1,2,3"))
    }

    @test fun `1,2,4 hash`() {
        assertEquals("63960835bcdc130f0b66d7ff4f6a5a8e" , hash("1,2,4"))
    }
}
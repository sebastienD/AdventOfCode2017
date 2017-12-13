import kotlin.test.assertEquals
import org.junit.Test as test

class Main_testKt() {

    @test fun f() {
        val lengths: List<Int> = mutableListOf(3, 4, 1, 5)

        val actual = launch(lengths, 5)

        assertEquals(listOf(2,1,0,3,4) , actual)
    }
}
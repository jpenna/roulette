import { getExpectedFor } from "./expectations";
import { last3 } from "./last";

export const rouletteNumbers: number[] = [
    10, 5, 24, 16, 33, 1, 20, 14, 31, 9, 22, 18, 29, 7, 28, 12, 35,
    3, 26, 0, 32, 15, 19, 4, 21, 2, 25, 17, 34, 6, 27, 13, 36,
    11, 30, 8, 23
];


export const rouletteNumberToIndex: Map<number, number> = new Map(
    rouletteNumbers.map((number, index) => [number, index])
);

export function compareWithPrevious(drawn: number, game: number[]) {



    for (const expected of game) {
        const index = rouletteNumberToIndex.get(expected) as number

        const prev = rouletteNumbers[index - 1 < 0 ? rouletteNumbers.length - 1 : index - 1]
        const target = rouletteNumbers[index]
        const next = rouletteNumbers[index + 1 > rouletteNumbers.length - 1 ? 0 : index + 1]

        if (prev === drawn || target === drawn || next === drawn) {
            return true
        }
    }
    return false

}

export function play(numbers: number[], protection: number = 0) {
    let curGame: number[] = []
    let wonPrev = true
    let usedCount = 0

    return numbers.slice(0, -1) // Remove last number
        .map((drawn, i) => {
            const previous = numbers[i + 1];

            curGame = wonPrev || usedCount >= protection ? getExpectedFor(previous) : curGame

            const result = compareWithPrevious(drawn, curGame)

            if (!result) {
                usedCount++
                wonPrev = false
                return false
            }

            usedCount = 0
            wonPrev = true
            return true
        });
}

function getMoney(numbers: number[], bet: number = 1, protection: number = 0, martin: boolean = false) {
    const results = play(numbers, protection);
    let result = 0;
    let pow = 0;

    let worst = 0

    let lastLostIndex = -1
    let lastLostSeq = 0
    let worstLoss = -1;

    let lostSeq = 0


    const total = numbers.length
    const matches = results.reduce((acc, win, index) => {
        if (win) {
            result += bet * 36 * (2 ** pow)
        } else {
            result -= bet * 12 * (2 ** pow)
        }

        if (result < worst) {
            worst = result
            worstLoss = index
        }

        if (win) {
            lastLostSeq = 0
            pow = 0
        } else {
            if (index === lastLostIndex + 1) {
                lastLostSeq++
            }

            lastLostIndex = index

            if (martin) {
                pow++
            }
        }

        if (lastLostSeq > lostSeq) {
            lostSeq = lastLostSeq
        }

        return win ? acc + 1 : acc
    }, 0)

    console.log(`Out of ${total} numbers, ${matches} numbers were found to be in the expected range`)

    console.log(`You would have made R$ ${result.toFixed(2)}`)
    console.log(`Worst case scenario: R$ ${worst.toFixed(2)}`)
    console.log(`Longest lost sequence: ${lostSeq}`)
    console.log(`Worst loss at index: ${worstLoss}`)
    console.log(`---\n`)
}

// getMoney(last3.slice(0, 10), 1, 2)
// getMoney(last3.slice(10, 20), 1, 2)
// getMoney(last3.slice(20, 30), 1, 2)
// getMoney(last3.slice(30, 40), 1, 2)
// getMoney(last3.slice(40, 50), 1, 2)
// getMoney(last3.slice(50, 60), 1, 2)
// getMoney(last3.slice(60, 70), 1, 2)
// getMoney(last3.slice(70, 80), 1, 2)
// getMoney(last3.slice(80, 90), 1, 2)
// getMoney(last3.slice(90, 100), 1, 2)
// getMoney(last3.slice(100, 110), 1, 2)
// getMoney(last3.slice(110, 120), 1, 2)
// getMoney(last3.slice(120, 130), 1, 2)
// getMoney(last3.slice(130, 140), 1, 2)
// getMoney(last3.slice(140, 150), 1, 2)
// getMoney(last3.slice(150, 160), 1, 2)
// getMoney(last3.slice(160, 170), 1, 2)
// getMoney(last3.slice(170, 180), 1, 2)
// getMoney(last3.slice(180, 190), 1, 2)
// getMoney(last3.slice(190, 200), 1, 2)
// getMoney(last3.slice(200, 210), 1, 2)
// getMoney(last3.slice(210, 220), 1, 2)
// getMoney(last3.slice(220, 230), 1, 2)
// getMoney(last3.slice(230, 240), 1, 2)
// getMoney(last3.slice(240, 250), 1, 2)
// getMoney(last3.slice(250, 260), 1, 2)
// getMoney(last3.slice(260, 270), 1, 2)
// getMoney(last3.slice(270, 280), 1, 2)
// getMoney(last3.slice(280, 290), 1, 2)
// getMoney(last3.slice(290, 300), 1, 2)
// getMoney(last3.slice(300, 310), 1, 2)
// getMoney(last3.slice(310, 320), 1, 2)
// getMoney(last3.slice(320, 330), 1, 2)
// getMoney(last3.slice(330, 340), 1, 2)
// getMoney(last3.slice(340, 350), 1, 2)
// getMoney(last3.slice(350, 360), 1, 2)
// getMoney(last3.slice(360, 370), 1, 2)
// getMoney(last3.slice(370, 380), 1, 2)
// getMoney(last3.slice(380, 390), 1, 2)
// getMoney(last3.slice(390, 400), 1, 2)
// getMoney(last3.slice(400, 410), 1, 2)
// getMoney(last3.slice(410, 420), 1, 2)
// getMoney(last3.slice(420, 430), 1, 2)
// getMoney(last3.slice(430, 440), 1, 2)
// getMoney(last3.slice(440, 450), 1, 2)
// getMoney(last3.slice(450, 460), 1, 2)
// getMoney(last3.slice(460, 470), 1, 2)
// getMoney(last3.slice(470, 480), 1, 2)
// getMoney(last3.slice(480, 490), 1, 2)
// getMoney(last3.slice(490, 500), 1, 2)

getMoney(last3, 2.5, 2)
getMoney(last3, 2.5, 2, true)
// getMoney(diffLast, 5, 2)
// getMoney(diffLast, 5, 2, true)



async function readAndPrint() {
    let gameNumber = 0

    let winCount = 0
    let lostCount = 0

    let lostSeq = 0

    while (true) {
        console.log(`Game ${gameNumber}`)
        const input = await prompt("Enter a number: ") as string
        gameNumber++
        const num = parseInt(input)
        if (!isNaN(num)) {
            const expected = getExpectedFor(num)
            const sorted = expected.sort((a, b) => rouletteNumberToIndex.get(a) as number - (rouletteNumberToIndex.get(b) as number))
            const expectedStr = sorted.map(e => `${e} (position: ${(rouletteNumberToIndex.get(e) as number) + 1})`)
            console.log(`Expected numbers for ${num}:\n\n`, expectedStr.join("\n "), "\n\n")
        } else {
            console.log("Please enter a valid number")
        }
    }
}

readAndPrint()

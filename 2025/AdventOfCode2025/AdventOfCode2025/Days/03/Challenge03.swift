import Foundation

struct Data03 {
    var nums: [Int]
}

struct Challenge03: Challenge, Identifiable {
    var id = UUID()
    var day: String { "03" }
    var description: String { """
        Part 1: Find the largest 2 diget numbers from left to right. \n
        Part 2: Find the largest 12 diget number from left to right.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "03")) ?? ""
    }
    
    private func parseInput() throws -> ([Data03]) {
        var data: [Data03] = []

        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            if line.isEmpty {
                continue
            }
            var nums: [Int] = []
            for num in line {
                nums.append(Int(String(num))!)
            }
            data.append(Data03(nums: nums))
        }
        
        return data
    }

    func part1() async throws -> String {
        let data = try parseInput()
        
        var result = 0
        for bank in data {
            // find max number that is not at the very end
            var firstHigh = 0
            for i in 1 ... bank.nums.count-2 {
                if bank.nums[i] > bank.nums[firstHigh] {
                    firstHigh = i
                }
            }
            // find the 2nd largest number after the max
            var nextHigh = firstHigh+1
            for i in firstHigh+1 ... bank.nums.count-1 {
                if bank.nums[i] > bank.nums[nextHigh] {
                    nextHigh = i
                }
            }
            result += Int(String(bank.nums[firstHigh]) + String(bank.nums[nextHigh]))!
        }
        
        return String(result)
    }

    func part2() async throws -> String {
        let data = try parseInput()
        let max = 12
        
        var result = 0
        for bank in data {
            var nums: [Int] = []
            var lastNum = 0
            for size in (1...max).reversed() {
                var maxNum = lastNum
                for i in lastNum...bank.nums.count-size {
                    if bank.nums[i] > bank.nums[maxNum] {
                        maxNum = i
                    }
                }
                nums.append(bank.nums[maxNum])
                lastNum = maxNum + 1
            }
            result += Int(nums.map { String($0) }.joined())!
        }
        
        return String(result)
    }
}

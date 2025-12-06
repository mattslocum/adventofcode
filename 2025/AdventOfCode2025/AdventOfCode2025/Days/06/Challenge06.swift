import Foundation

struct Data06 {
    var nums: [Int]
    var op: String
}

struct Challenge06: Challenge, Identifiable {
    var id = UUID()
    var day: String { "06" }
    var description: String { """
        Part 1: Use the operator in the vertical column to calculate the result. \n
        Part 2: Use the operator in the vertical column, but the numbers are also stacked vertically.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "06")) ?? ""
    }
    
    private func parseInput() throws -> ([Data06]) {
        let lines = input.components(separatedBy: CharacterSet.newlines)
        let linesSplit = lines.map({ line in
            return line.components(separatedBy: CharacterSet.whitespaces).filter { !$0.isEmpty }
        })
        var data: [Data06] = []
        for i in 0..<linesSplit[0].count {
            let nums = linesSplit.map({ Int($0[i]) }).filter { $0 != nil }
            let op = linesSplit[lines.count-1][i]
            data.append(Data06(nums: nums as! [Int], op: op))
        }
        return data
    }

    func part1() async throws -> String {
        let list = try parseInput()
        
        var result = 0
        for item in list {
            switch item.op {
            case "+":
                result += item.nums.reduce(0, +)
            case "*":
                result += item.nums.reduce(1, *)
            default:
                print("Invalid operator: \(item.op)")
            }
        }
        
        return String(result)
    }

    private func parseInput2() throws -> ([Data06]) {
        let linesRaw = input.components(separatedBy: CharacterSet.newlines)
        let lines = linesRaw.map { Array($0) }
        var sets: [(Int, Int)] = []
        var lastSymbol = 0
        for i in 1..<lines.last!.count {
            if lines.last![i] == "+" || lines.last![i] == "*" {
                sets.append((lastSymbol, i-2))
                lastSymbol = i
            }
        }
        sets.append((lastSymbol, linesRaw[0].count-1))

        var data: [Data06] = []
        for set in sets {
            var nums: [Int] = []
            for i in set.0...set.1 {
                var num = ""
                for y in 0..<lines.count-1 {
                    if i < lines[y].count && lines[y][i] != " " {
                        num += String(lines[y][i])
                    }
                }
                print(num)
                nums.append(Int(num)!)
            }
            data.append(Data06(nums: nums, op: String(lines.last![set.0])))
        }
        return data
    }

    func part2() async throws -> String {
        let list = try parseInput2()
        
        var result = 0
        for item in list {
            switch item.op {
            case "+":
                result += item.nums.reduce(0, +)
            case "*":
                result += item.nums.reduce(1, *)
            default:
                print("Invalid operator: \(item.op)")
            }
        }
        
        return String(result)
    }
}

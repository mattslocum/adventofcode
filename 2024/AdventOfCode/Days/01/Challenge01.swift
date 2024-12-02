import Foundation

struct Challenge01: Challenge, Identifiable {
    var id = UUID()
    var day: String { "01" }
    var description: String { """
        Part 1: Calculate the distance between two points. \n
        Part 2: Calculate how many times the first number occures in the 2nd list then multiplied by that number.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "01")) ?? ""
    }
    
    private func parseInput() throws -> ([Int], [Int]) {
        var first: [Int] = []
        var second: [Int] = []

        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            if line.isEmpty {
                continue
            }
            let twoNum = line.components(separatedBy: "   ")
            first.append(Int(twoNum[0])!)
            second.append(Int(twoNum[1])!)
        }
        
        return (first, second)
    }

    func part1() async throws -> String {
        var (first, second) = try parseInput()
        first.sort()
        second.sort()
        
        var result = 0
        for (i, val) in first.enumerated() {
            result += abs(val - second[i])
        }
        
        return String(result)
    }

    func part2() async throws -> String {
        let (first, second) = try parseInput()

        var counts = [Int:Int]()
        for val in second {
            counts[val, default: 0] += 1
        }
        
        var result = 0
        for val in first {
            result += val * counts[val, default: 0]
        }

        return String(result)
    }
}

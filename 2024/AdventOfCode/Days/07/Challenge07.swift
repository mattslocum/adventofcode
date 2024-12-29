import Foundation
import RegexBuilder

struct Data07 {
    var answer: Int
    var params: [Int]
}

struct Challenge07: Challenge, Identifiable {
    var id = UUID()
    var day: String { "07" }
    var description: String { """
        Part 1: Found how many numbers can be calculated (left to right) by either using a + or *.\n
        Part 2: Same as above but also allows for concatenation.
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "07")) ?? ""
    }

    private func parseInput() throws -> ([Data07]) {
        var data: [Data07] = []
        
        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            let parts = line.components(separatedBy: ":")
            let answer = Int(parts[0])!
            let params = parts[1].trimmingCharacters(in: .whitespaces).components(separatedBy: " ").map({ Int($0)! })
            data.append(Data07(answer: answer, params: params))
        }
        return data
    }

    func part1() async throws -> String {
        let data = try parseInput()
        var answers = 0
        
        for d in data {
            var tries: [Int] = [d.params[0]]
            for i in 1..<d.params.count {
                let mul = tries.map({ $0 * d.params[i] })
                let add = tries.map({ $0 + d.params[i] })
                tries = mul + add
            }
            if tries.contains(d.answer) {
                answers += d.answer
            }
        }
        
        return String(answers)
    }
    
    func part2() async throws -> String {
        let data = try parseInput()
        var answers = 0
        
        for d in data {
            var tries: [Int] = [d.params[0]]
            for i in 1..<d.params.count {
                let mul = tries.map({ $0 * d.params[i] })
                let add = tries.map({ $0 + d.params[i] })
                let concat = tries.map({ Int(String($0) + String(d.params[i]))! })
                tries = mul + add + concat
            }
            if tries.contains(d.answer) {
                answers += d.answer
            }
        }
        
        return String(answers)
    }
}




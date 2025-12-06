import Foundation

struct Data01 {
    var dir: Character
    var amt: Int
}

struct Challenge01: Challenge, Identifiable {
    var id = UUID()
    var day: String { "01" }
    var description: String { """
        Part 1: Calculate the number of times the dial moves to zero. \n
        Part 2: Calculate the number of times we pass or land on zero.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "01")) ?? ""
    }
    
    private func parseInput() throws -> ([Data01]) {
        var data: [Data01] = []

        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            if line.isEmpty {
                continue
            }
            if let firstChar = line.first {
                let num = Int(String(line.dropFirst()))!
                data.append(Data01(dir: firstChar, amt: num))
            }
        }
        
        return data
    }

    func part1() async throws -> String {
        let data = try parseInput()
        var pos = 50
        let max = 100;
        
        var result = 0
        for d in data {
            switch d.dir {
                case "R":
                    pos += d.amt
                case "L":
                    pos -= d.amt
                default:
                    print("Invalid direction: \(d.dir)")
            }
            pos %= max;
            if pos == 0 {
                result += 1
            }
        }
        
        return String(result)
    }

    func part2() async throws -> String {
        let data = try parseInput()
        var pos = 50
        let max = 100;
        
        var result = 0
        for d in data {
            switch d.dir {
                case "R":
                    result += Int(d.amt / max)
                    let move = d.amt % max
                    if pos + move > 99 {
                        result += 1
                        pos = (pos + move) % max
                    } else {
                        pos += move
                    }
                case "L":
                    result += Int(d.amt / max)
                    let move = d.amt % max
                    if pos == 0 {
                        pos = max - move
                    } else if pos - move <= 0 {
                        result += 1
                        pos = (max + (pos - move)) % max
                    } else {
                        pos -= move
                    }
                default:
                    print("Invalid direction: \(d.dir)")
            }
        }
        
        return String(result)
    }
}

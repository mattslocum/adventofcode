import Foundation

struct Data05 {
    var ranges: [(Int, Int)]
    var ingredients: [Int]
}

struct Challenge05: Challenge, Identifiable {
    var id = UUID()
    var day: String { "05" }
    var description: String { """
        Part 1:  \n
        Part 2: 
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "05")) ?? ""
    }
    
    private func parseInput() throws -> Data05 {
        let lines = input.components(separatedBy: CharacterSet.newlines)
        var ranges: [(Int, Int)] = []
        var ingredients: [Int] = []
        var isDoneRanges = false
        for line in lines {
            if line == "" {
                isDoneRanges = true
                continue
            }
            if !isDoneRanges {
                let parts = line.components(separatedBy: "-")
                ranges.append((Int(parts[0])!, Int(parts[1])!))
            } else {
                ingredients.append(Int(line)!)
            }
        }
        // sort ranges by start
        ranges.sort { $0.0 < $1.0 }

        return Data05(ranges: ranges, ingredients: ingredients)
    }

    func part1() async throws -> String {
        let data = try parseInput()

        var result = 0
        for ingredient in data.ingredients {
            for range in data.ranges {
                if ingredient >= range.0 && ingredient <= range.1 {
                    result += 1
                    break
                } else if ingredient < range.0 {
                    break
                }
            }
        }

        return String(result)
    }

    func part2() async throws -> String {
        var data = try parseInput()
        
        var ranges: [(Int, Int)] = []
        // merge ranges. Note: they are already sorted by start
        for range in data.ranges {
            if let lastRange = ranges.last, lastRange.1 >= range.0 {
                ranges[ranges.count-1] = (lastRange.0, max(lastRange.1, range.1))
            } else {
                ranges.append(range)
            }
        }

        var result = 0
        for range in ranges {
            result += range.1 - range.0 + 1 // +1 because inclusive
        }

        return String(result)
    }
}

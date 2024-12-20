import Foundation
import RegexBuilder

struct Challenge09: Challenge, Identifiable {
    var id = UUID()
    var day: String { "09" }
    var description: String { """
        Part 1: \n
        Part 2: 
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "09")) ?? ""
    }

    private func parseInput() throws -> ([Int], [Int]) {
        var values: [Int] = []
        var gaps: [Int] = []
        var isVal = true
        for val in input {
            if isVal {
                // TODO: Better way for char to Int
                values.append(Int(String(val))!)
            } else {
                // TODO: Better way for char to Int
                gaps.append(Int(String(val))!)
            }
            isVal = !isVal
        }
        return (values, gaps)
    }

    // 6367087064415 is the answer
    func part1() async throws -> String {
        var (values, gaps) = try parseInput()
        let first = values.removeFirst()
        var defrag: [Int] = Array(repeating: 0, count: first)
        for var (i, gap) in gaps.enumerated() {
            if i >= values.count {
                break
            }
            while gap > 0 {
                var lastVal = values[values.count - 1]
                if lastVal <= gap {
                    // not enough or equal to close the gap
                    defrag += Array(repeating: values.count, count: lastVal)
                    gap -= lastVal
                    lastVal = values.removeLast()
                    gaps.removeLast()
                } else {
                    // too many on the right side
                    defrag += Array(repeating: values.count, count: gap)
                    lastVal -= gap
                    // save how many we have left
                    values[values.count - 1] = lastVal
                    gap = 0
                }
                if i >= values.count {
                    break
                }
            }
            if i < values.count {
                defrag += Array(repeating: i+1, count: values[i])
            } else {
                print("out of range")
            }
        }

        var answer = 0
        for (i, val) in defrag.enumerated() {
            answer += i * val
        }

        return String(answer)
    }

    // Sample data: 2333133129494939491
    func part2() async throws -> String {
        var (values, gaps) = try parseInput()
        let first = values.removeFirst()
        var defrag: [Int] = Array(repeating: 0, count: first)
        var valsUsed: [Int] = []
        // 00...111...2...333.44.5555.6666.777.888899
        // 00992111777.44.333....5555.6666.....8888..
        // 130313103
        // 0......2...3...444
        // 0444322

        for var (i, gap) in gaps.enumerated() {
            // 1. find last values that can fill the gap
            for valIdx in (i..<values.count).reversed() {
                if valsUsed.contains(valIdx) { continue }
                let val = values[valIdx]
                if val <= gap {
                    defrag += Array(repeating: valIdx+1, count: val)
                    valsUsed.append(valIdx)
                    gap -= val
                }
                if gap == 0 {
                    break;
                }
            }
            if gap > 0 {
                // 2. empty fill
                defrag += Array(repeating: 0, count: gap)
            }
            // 3. add the next value
            if values.count > i {
                if !valsUsed.contains(i) {
                    // +1 because we already removed the first value before starting
                    defrag += Array(repeating: i+1, count: values[i])
                } else {
                    // moved val now creates a gap that can't be filled
                    defrag += Array(repeating: 0, count: values[i])
                }
            }
        }

        var answer = 0
        for (i, val) in defrag.enumerated() {
            answer += i * val
        }

        return String(answer)
    }
}




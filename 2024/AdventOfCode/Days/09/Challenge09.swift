import Foundation
import RegexBuilder

struct Challenge09: Challenge, Identifiable {
    var id = UUID()
    var day: String { "09" }
    var description: String { """
        Part 1: Defragment empty spaces by moving letters from the right side into the block on the left side.\n
        Part 2: Defragment empty spaces by moving whole letter sets instead.
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
        var valsUsed: [Int:Bool] = [:]
        // 00...111...2...333.44.5555.6666.777.888899
        // 00992111777.44.333....5555.6666.....8888..
        // 130313103
        // 0......2...3...444
        // 0444322

        for var (i, gap) in gaps.enumerated() {
            // 1. find last values that can fill the gap
            for valIdx in (i..<values.count).reversed() {
                if valsUsed[valIdx, default: false] { continue }
                let val = values[valIdx]
                if val <= gap {
                    defrag += Array(repeating: valIdx+1, count: val)
                    valsUsed[valIdx] = true
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
                if !valsUsed[i, default: false] {
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




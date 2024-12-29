import Foundation
import RegexBuilder

var cache11: [String: Int] = [:]
/**
 Initial arrangement:
 125 17

 After 1 blink:
 253000 1 7

 After 2 blinks:
 253 0 2024 14168

 After 3 blinks:
 512072 1 20 24 28676032

 After 4 blinks:
 512 72 2024 2 0 2 4 2867 6032

 After 5 blinks:
 1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32

 After 6 blinks:
 2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
 */

struct Challenge11: Challenge, Identifiable {
    var id = UUID()
    var day: String { "11" }
    var description: String { """
        Part 1: Tansform a set of numbers into other numbers based on a formula. Run that formula 25 times to get the count of the result.\n
        Part 2: Run the formula 75 times.
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "11")) ?? ""
    }

    private func parseInput() throws -> ([Int]) {
        return input.components(separatedBy: " ").map{ Int($0)! }
    }

    func part1() async throws -> String {
        var nums = try parseInput()
        
        for _ in 0..<25 {
            nums = runMutation(nums: nums)
            print(nums.count)
        }

        return String(nums.count)
    }
    
    func runMutation(nums: [Int]) -> [Int] {
        var list: [Int] = []
        for val in nums {
            let digits = String(val).count
            if val == 0 {
                list.append(1)
            } else if digits % 2 == 0 {
                let digitSplitter = pow(base: 10, pow: digits / 2)
                // first half of the digits.
                list.append(val / digitSplitter)
                list.append(val % digitSplitter)
            } else {
                list.append(val * 2024)
            }
        }

        return list
    }
    
    // TODO: replace with native
    private func pow(base: Int, pow: Int) -> Int {
        var answer = 1
        for _ in 0..<pow {
            answer *= base
        }
        return answer
    }

    func part2() async throws -> String {
        var nums = try parseInput()

        var found: Int = 0
        for n in nums {
            found += recurse(num: n, depth: 75)
        }

        return String(found)
    }
    
    func recurse(num: Int, depth: Int) -> Int {
        let cacheKey = String(num) + ":" + String(depth)
        if cache11[cacheKey] == nil {
            var list: Int = 0
            let digits = String(num).count
            if num == 0 {
                if depth == 1 {
                    list += 1
                } else {
                    list += recurse(num: 1, depth: depth - 1)
                }
            } else if digits % 2 == 0 {
                if depth == 1 {
                    list += 2
                } else {
                    let digitSplitter = pow(base: 10, pow: digits / 2)
                    // first half of the digits.
                    list += recurse(num: num / digitSplitter, depth: depth - 1)
                    list += recurse(num: num % digitSplitter, depth: depth - 1)
                }
            } else {
                if depth == 1 {
                    list += 1
                } else {
                    list += recurse(num: num * 2024, depth: depth - 1)
                }
            }
            cache11[cacheKey] = list
        }
        return cache11[cacheKey]!
    }
}




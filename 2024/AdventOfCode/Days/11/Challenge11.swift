import Foundation
import RegexBuilder

struct Challenge11: Challenge, Identifiable {
    var id = UUID()
    var day: String { "11" }
    var description: String { """
        Part 1: \n
        Part 2: 
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
                // first half of the digist.
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
        
        return "part 2"
    }
}




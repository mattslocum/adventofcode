import Foundation

enum Direction {
    case up
    case down
}

struct Challenge02: Challenge, Identifiable {
    var id = UUID()
    var day: String { "02" }
    var description: String { """
        Part 1: All numbers must be increasing or decreasing by 1-3. \n
        Part 2: One number can be removed to make the list good.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "02")) ?? ""
    }
    
    private func parseInput() throws -> ([[Int]]) {
        let lines = input.components(separatedBy: CharacterSet.newlines)
        return lines.map({ line in
            return line.components(separatedBy: " ").map({ Int($0)! })
        })
    }

    func part1() async throws -> String {
        let sets = try parseInput()
        var found = 0

        for nums in sets {
            let good = checkGood(nums: nums)
            if good {
                found += 1
            }
        }

        return String(found)
    }

    func part2() async throws -> String {
        let sets = try parseInput()
        var found = 0

        for nums in sets {
            var good = checkGood(nums: nums)
            if good {
                found += 1
                continue
            }
            for i in 0..<nums.count {
                var attempt = nums
                attempt.remove(at: i)
                good = checkGood(nums: attempt)
                if good {
                    found += 1
                    break
                }
            }
        }

        return String(found)
    }
    
    private func checkGood(nums: [Int]) -> Bool {
        let dir = nums[0] < nums[1] ? Direction.up : Direction.down
        var prev = nums[0]
        var good = true
        for i in 1..<nums.count {
            if prev == nums[i] {
                good = false
                break
            }
            if abs(prev - nums[i]) > 3 {
                good = false
                break
            }
            let curDir = prev < nums[i] ? Direction.up : Direction.down
            if curDir != dir {
                good = false
                break
            }
            prev = nums[i]
        }
        return good
    }
}

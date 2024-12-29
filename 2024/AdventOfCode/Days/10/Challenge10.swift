import Foundation
import RegexBuilder

struct Challenge10: Challenge, Identifiable {
    var id = UUID()
    var day: String { "10" }
    var description: String { """
        Part 1: Count the number of 9s that can be reached from starting at 0 and counting up.\n
        Part 2: Count the number of unique paths from 0 to 9.
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "10")) ?? ""
    }

    // first return is the map, 2nd is the list of y,x 0s
    private func parseInput() throws -> ([[Int]], [[Int]]) {
        var map: [[Int]] = []
        var zeros: [[Int]] = []
        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            var nums: [Int] = []
            for char in line {
                let val = Int(String(char)) ?? -1
                if val == 0 {
                    zeros.append([map.count, nums.count])
                }
                nums.append(val)
            }
            map.append(nums)
        }
        return (map, zeros)
    }

    func part1() async throws -> String {
        let (map, zeros) = try parseInput()
        var result = 0
        for start in zeros {
            let (val, _) = findEnds(map: map, startYX: start)
            result += val
        }

        return String(result)
    }
    
    func findEnds(map: [[Int]], startYX: [Int]) -> (Int, Int) {
        var finishes: [[Int]] = []
        var pathsYX = [startYX]
        var totalTrails = 1
        // start tree walking. Depth first search
        while pathsYX.count > 0 {
            let curYX = pathsYX.removeFirst()
            let nextVal = map[curYX[0]][curYX[1]] + 1
            if nextVal == 10 {
                // We are already done. Save it.
                if !finishes.contains(curYX) {
                    finishes.append(curYX)
                }
                continue
            }
            var origPaths = pathsYX
            // check neighbors
            // up
            if curYX[0] > 0 && map[curYX[0]-1][curYX[1]] == nextVal {
                pathsYX.append([curYX[0]-1, curYX[1]])
            }
            // down
            if curYX[0] + 1 < map.count && map[curYX[0]+1][curYX[1]] == nextVal {
                pathsYX.append([curYX[0]+1, curYX[1]])
            }
            // left
            if curYX[1] > 0 && map[curYX[0]][curYX[1]-1] == nextVal {
                pathsYX.append([curYX[0], curYX[1]-1])
            }
            // right
            if curYX[1] + 1 < map[0].count && map[curYX[0]][curYX[1]+1] == nextVal {
                pathsYX.append([curYX[0], curYX[1]+1])
            }
            totalTrails += pathsYX.count - origPaths.count - 1
        }
        print(startYX, "|", finishes)
        
        return (finishes.count, totalTrails)
    }

    func part2() async throws -> String {
        let (map, zeros) = try parseInput()
        var result = 0
        for start in zeros {
            let (_, paths) = findEnds(map: map, startYX: start)
            result += paths
        }

        return String(result)
    }
}




import Foundation

struct Challenge04: Challenge, Identifiable {
    var id = UUID()
    var day: String { "04" }
    var description: String { """
        Part 1: Find how many @ have < 4 adjacent @s. \n
        Part 2: Find how many @ have < 4 adjacent, but remove until all can be removed.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "04")) ?? ""
    }
    
    private func parseInput() throws -> ([[String]]) {
        let lines = input.components(separatedBy: CharacterSet.newlines)
        return lines.map({ line in
            return Array(line).map({ String($0) })
        })
    }

    func part1() async throws -> String {
        let grid = try parseInput()
        
        var result = 0
        for y in 0..<grid.count {
            for x in 0..<grid[y].count {
                if grid[y][x] != "@" {
                    continue
                }
                var neighbors = 0
                for dirY in -1...1 {
                    if y + dirY < 0 || y + dirY >= grid.count {
                        continue
                    }
                    for dirX in -1...1 {
                        if x + dirX < 0 || x + dirX >= grid[y].count {
                            continue
                        }
                        if dirY == 0 && dirX == 0 {
                            continue
                        }
                        if grid[y+dirY][x+dirX] == "@" {
                            neighbors += 1
                        }
                    }
                }
                if neighbors < 4 {
                    result += 1
                }
            }
        }
        
        return String(result)
    }

    func part2() async throws -> String {
        var grid = try parseInput()
        
        var result = 0
        while true {
            let sparseCount = countSparseAtSymbols(in: &grid)
            if sparseCount == 0 {
                break
            }
            result += sparseCount
        }
        printGrid(grid)

        return String(result)
    }

    private func countSparseAtSymbols(in grid: inout [[String]]) -> Int {
        var result = 0
        for y in 0..<grid.count {
            for x in 0..<grid[y].count {
                if grid[y][x] != "@" {
                    continue
                }
                var neighbors = 0
                for dirY in -1...1 {
                    if y + dirY < 0 || y + dirY >= grid.count {
                        continue
                    }
                    for dirX in -1...1 {
                        if x + dirX < 0 || x + dirX >= grid[y].count {
                            continue
                        }
                        if dirY == 0 && dirX == 0 {
                            continue
                        }
                        if grid[y+dirY][x+dirX] == "@" {
                            neighbors += 1
                        }
                    }
                }
                if neighbors < 4 {
                    grid[y][x] = "x"
                    result += 1
                }
            }
        }
        return result
    }

    private func printGrid(_ grid: [[String]]) {
        for row in grid {
            print(row.joined())
        }
    }
}

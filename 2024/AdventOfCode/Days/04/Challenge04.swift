import Foundation
import RegexBuilder


struct Challenge04: Challenge, Identifiable {
    var id = UUID()
    var day: String { "04" }
    var description: String { """
        Part 1: \n
        Part 2: 
        """ }
    var input: String
    var xmas = Array("XMAS").map({ String($0) })
    var mas = Array("MAS").map({ String($0) })

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
        var found = 0

        for y in 0..<grid.count {
            for x in 0..<grid[y].count {
                if grid[y][x] == xmas[0] {
                    found += numXMAS(grid: grid, x: x, y: y)
                }
            }
        }

        return String(found)
    }
    
    private func numXMAS(grid: [[String]], x: Int, y: Int) -> Int {
        var found = 0
        for dirY in -1...1 {
            for dirX in -1...1 {
                if dirY == 0 && dirX == 0 {
                    continue
                }
                for i in 1...xmas.count-1 {
                    let xchar = xmas[i]
                    let adjY = y + (dirY * (i))
                    let adjX = x + (dirX * (i))
                    if adjY < 0 || adjY >= grid.count || adjX < 0 || adjX >= grid[0].count {
                        break
                    }
                    if grid[adjY][adjX] != xchar {
                        break
                    }
                    if i == xmas.count - 1 {
                        found += 1
                    }
                }
            }
        }
        return found
    }

    func part2() async throws -> String {
        let grid = try parseInput()
        var found = 0

        for y in 1..<grid.count-1 {
            for x in 1..<grid[y].count-1 {
                if grid[y][x] == "A" && hasCrossMAS(grid: grid, x: x, y: y) {
                    found += 1
                }
            }
        }

        return String(found)
    }
    
    private func hasCrossMAS(grid: [[String]], x: Int, y: Int) -> Bool {
        var found = 0
        var chars = [String]()
        
        // get corner characters
        for dirY in -1...1 {
            if dirY == 0 { continue }
            for dirX in -1...1 {
                if dirX == 0 { continue }
                
                let adjY = y + dirY
                let adjX = x + dirX
                chars.append(grid[adjY][adjX])
            }
        }
        
        if (chars[0] == "M" && chars[3] == "S") ||
            (chars[0] == "S" && chars[3] == "M")
        {
            found += 1
        }
        if (chars[1] == "M" && chars[2] == "S") ||
            (chars[1] == "S" && chars[2] == "M")
        {
            found += 1
        }

        return found == 2
    }
    
//    private getChars
}




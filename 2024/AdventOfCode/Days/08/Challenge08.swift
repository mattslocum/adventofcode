import Foundation
import RegexBuilder

struct Challenge08: Challenge, Identifiable {
    var id = UUID()
    var day: String { "08" }
    var description: String { """
        Part 1: \n
        Part 2: 
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "08")) ?? ""
    }

    private func parseInput() throws -> ([Character: [[Int]]], Int, Int) {
        var data: [Character: [[Int]]] = [:]
        
        let lines = input.components(separatedBy: CharacterSet.newlines)
        for (y, line) in lines.enumerated() {
            for (x, char) in line.enumerated() {
                if char == "." {
                    continue
                }
                data[char] = data[char, default: []] + [[y, x]]
            }
        }
        return (data, lines.count, lines[0].count)
    }

    func part1() async throws -> String {
        let (data, maxY, maxX) = try parseInput()
        var found: [[Int]] = []
        
        for nodes in data.values {
            for i in 0..<nodes.count {
                for j in 0..<nodes.count {
                    if i == j {
                        continue
                    }
                    let y = nodes[i][0] + nodes[i][0] - nodes[j][0]
                    let x = nodes[i][1] + nodes[i][1] - nodes[j][1]
                    if y < 0 || x < 0 || y >= maxY || x >= maxX {
                        continue
                    }
                    if !found.contains([y, x]) {
                        found.append([y, x])
                    }
                }
            }
        }
        
        return String(found.count)
    }
    
    // 1075 is too low
    func part2() async throws -> String {
        let (data, maxY, maxX) = try parseInput()
        var found: [[Int]] = []
        
        for nodes in data.values {
            for i in 0..<nodes.count {
                for j in 0..<nodes.count {
                    if i == j {
                        continue
                    }
                    let offY = nodes[i][0] - nodes[j][0]
                    let offX = nodes[i][1] - nodes[j][1]
                    var y = nodes[i][0]
                    var x = nodes[i][1]
                    while y >= 0 && x >= 0 && y < maxY && x < maxX {
                        if !found.contains([y, x]) {
                            found.append([y, x])
                        }
                        y += offY
                        x += offX
                    }
                }
            }
        }
        
        return String(found.count)
    }
}




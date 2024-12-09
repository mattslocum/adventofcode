import Foundation
import RegexBuilder


struct Challenge05: Challenge, Identifiable {
    var id = UUID()
    var day: String { "05" }
    var description: String { """
        Part 1: \n
        Part 2: 
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "05")) ?? ""
    }

    private func parseInput() throws -> ([Int: [Int]], [[Int]]) {
        var pagesDictionary: [Int: [Int]] = [:]
        var pageList: [[Int]] = []
        var foundSeparator = false
        
        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            if line == "" {
                foundSeparator = true
                continue
            }
            if !foundSeparator {
                let set = line.components(separatedBy: "|").map{ Int($0) }
                pagesDictionary[set[0]!, default: []].append(set[1]!)
            } else {
                let pages = (line.components(separatedBy: ",").map{ Int($0) } as? [Int])!
                pageList.append(pages)
            }
        }
        
        return (pagesDictionary, pageList)
    }

    func part1() async throws -> String {
        let (sets, lists) = try parseInput()
//        print(sets)
//        print(lists)
        var found = 0

        for list in lists {
            let (first, second) = invalidPages(sets: sets, list: list)
            if first == 0 && second == 0 {
//                print(list)
                let center = list.count/2
//                print("Center: \(list[center])")
                found += list[center]
            }
        }

        return String(found)
    }
    
    func invalidPages(sets: [Int: [Int]], list: [Int]) -> (Int, Int) {
        // start 1 in because we are checking for prior conflicts
        for i in 1..<list.count {
            let page = list[i]
            let mustBeAfter = sets[page]
            if mustBeAfter == nil {
                continue
            }
            for j in 0..<i {
                let prevPage = list[j]
                if mustBeAfter!.contains(prevPage) {
                    return (j, i)
                }
            }
        }
        return (0, 0)
    }

    func part2() async throws -> String {
        let (sets, lists) = try parseInput()
//        print(sets)
//        print(lists)
        var found = 0

        for var list in lists {
            var (first, second) = invalidPages(sets: sets, list: list)
            if first != 0 || second != 0 {
                while (first != 0 || second != 0) {
                    let tmpSecond = list[second]
                    list[second] = list[first]
                    list[first] = tmpSecond
                    (first, second) = invalidPages(sets: sets, list: list)
                }
//                print(list)
                let center = list.count/2
//                print("Center: \(list[center])")
                found += list[center]
            }
        }

        return String(found)
    }
}




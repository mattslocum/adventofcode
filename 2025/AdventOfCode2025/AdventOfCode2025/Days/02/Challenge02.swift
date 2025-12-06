import Foundation

struct Data02 {
    var start: Int
    var end: Int
}

struct Challenge02: Challenge, Identifiable {
    var id = UUID()
    var day: String { "02" }
    var description: String { """
        Part 1: Find numbers in each range that have duplicate numbers. \n
        Part 2: Find numbers in each range that have perfectly repeating numbers.
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "02")) ?? ""
    }
    
    private func parseInput() throws -> ([Data02]) {
        var data: [Data02] = []

        let items = input.components(separatedBy: ",")
        for item in items {
            if item.isEmpty {
                continue
            }
            let range = item.components(separatedBy: "-")
            if range.count != 2 {
                continue
            }
            data.append(Data02(start:Int(range[0])!, end:Int(range[1])!))
        }
        
        return data
    }

    func part1() async throws -> String {
        let data = try parseInput()
        
        var result = 0
        for d in data {
            for i in d.start...d.end {
                result += getRepeatingHalf(num: i)
            }
        }
        
        return String(result)
    }
    
    func getRepeatingHalf(num: Int) -> Int {
        let digits = String(num)
        if digits.count % 2 != 0 {
            return 0
        }
        
        // split string in half and check if the first half is the same as the second half
        let half = digits.count / 2
        let firstHalf = String(digits.prefix(half))
        let secondHalf = String(digits.suffix(half))
        if firstHalf == secondHalf {
            return num
        }
        
        return 0
    }

    func part2() async throws -> String {
        let data = try parseInput()
        
        var result = 0
        for d in data {
            for i in d.start...d.end {
                result += getRepeatingNum(num: i)
            }
        }

        return String(result)
    }
    
    func getRepeatingNum(num: Int) -> Int {
        let digits = String(num)
        if digits.count < 2 {
            return 0
        }
        
        for i in 1 ... digits.count/2 {
            // split the string by the size of i, and check if each piece is the same
            let pieces = split(str: digits, by: i)
            var equal = true
            for piece in pieces {
                if piece != pieces[0] {
                    equal = false
                    break
                }
            }
            if equal {
                return num
            }
        }
        
        return 0
    }
    
    func split(str: String, by: Int) -> [String] {
        let group = Array(str).chunked(into: by)
        return group.map { String($0) }
    }
}

extension Array {
    func chunked(into size: Int) -> [[Element]] {
        return stride(from: 0, to: count, by: size).map {
            Array(self[$0 ..< Swift.min($0 + size, count)])
        }
    }
}

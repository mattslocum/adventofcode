import Foundation

struct Challenge01: Challenge, Identifiable {
    var id = UUID()
    var day: String { "01" }
    
    private func parseInput() throws -> ([Int], [Int]) {
        guard let fileURL = Bundle.main.url(forResource: "input", withExtension: "txt") else {
            throw NSError(domain: "File not found", code: 404, userInfo: nil)
        }

        let contents = try String(contentsOf: fileURL, encoding: .utf8)
        var first: [Int] = []
        var second: [Int] = []

        let lines = contents.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            if line.isEmpty {
                continue
            }
            let twoNum = line.components(separatedBy: "   ")
            first.append(Int(twoNum[0])!)
            second.append(Int(twoNum[1])!)
        }
        
        return (first, second)
    }

    func run1() async throws -> String {
        var (first, second) = try parseInput()
        first.sort()
        second.sort()
        
        var result = 0
        for (i, val) in first.enumerated() {
            result += abs(val - second[i])
        }
        
        return String(result)
    }
    func run2() async throws -> String {
        let (first, second) = try parseInput()

        var counts = [Int:Int]()
        for val in second {
            counts[val, default: 0] += 1
        }
        
        var result = 0
        for val in first {
            result += val * counts[val, default: 0]
        }

        return String(result)
    }
}

import Foundation

struct Challenge02: Challenge, Identifiable {
    var id = UUID()
    var day: String { "02" }
    var description: String { """
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "02")) ?? ""
        print(input)
    }

    private func parseInput() throws -> String {
        return input
    }

    func part1() async throws -> String {
        // Challenge 1 logic
        return "Result of Challenge 2a"
    }
    func part2() async throws -> String {
        // Challenge 1 logic
        return "Result of Challenge 2b"
    }
}

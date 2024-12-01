import SwiftUI

struct ChallengeDetailView: View {
    let challenge: any Challenge
    @State private var result: String = ""

    var body: some View {
        VStack {
            HStack {
                Button("Run Part 1") {
                    Task {
                        do {
                            result = try await challenge.run1()
                        } catch {
                            result = "Error: \(error.localizedDescription)"
                        }
                    }
                }
                Button("Run Part 2") {
                    Task {
                        do {
                            result = try await challenge.run2()
                        } catch {
                            result = "Error: \(error.localizedDescription)"
                        }
                    }
                }
            }
            .padding()

            ScrollView {
                Text(result)
                    .font(.system(size: 14))
                    .padding()
                    .frame(maxWidth: .infinity, alignment: .topLeading)
                    .border(Color.gray, width: 1)
            }
        }
    }
}

import SwiftUI

// struct ChallengeDetailView<ChallengeType: Challenge & Equatable>: View {
//     let challenge: ChallengeType

struct ChallengeDetailView: View {
    let challenge: any Challenge
    @State private var result: String = ""
    @State private var inputCollapsed: Bool = true
    @State private var time: Double = 0.0

    var body: some View {
        VStack {
            Text(challenge.description)
                .font(.system(size: 14))
                .padding()
            
            HStack {
                Button("Run Part 1") {
                    Task {
                        let start = Date()
                        do {
                            result = try await challenge.part1()
                        } catch {
                            result = "Error: \(error.localizedDescription)"
                        }
                        time = Date().timeIntervalSince(start)
                    }
                }
                Button("Run Part 2") {
                    Task {
                        let start = Date()
                        do {
                            result = try await challenge.part2()
                        } catch {
                            result = "Error: \(error.localizedDescription)"
                        }
                        time = Date().timeIntervalSince(start)
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
            
            VStack {
                Button(
                    action: { self.inputCollapsed.toggle() },
                    label: {
                        HStack {
                            Text("Input Data")
                            Spacer()
                            Image(systemName: self.inputCollapsed ? "chevron.down" : "chevron.up")
                        }
                        .padding(.bottom, 1)
                        .background(Color.white.opacity(0.01))
                    }
                )
                .buttonStyle(PlainButtonStyle())
                
                ScrollView {
                    Text(self.challenge.input)
                        .frame(maxWidth: .infinity, alignment: .topLeading)
                }
                .padding()
                .frame(maxWidth: .infinity, maxHeight: inputCollapsed ? 0 : .none, alignment: .topLeading)
                .clipped()
                .animation(.easeOut, value: inputCollapsed)
                .transition(.slide)
            }
            
            Text("Execution time: \(time)")
                .font(.system(size: 12))
                .frame(maxWidth: .infinity, alignment: .topLeading)
        }
        .onChange(of: challenge.id) { /* newChallenge in */
            // Reset the private variable when the challenge changes
            result.self = ""
            inputCollapsed.self = true
        }
    }
}

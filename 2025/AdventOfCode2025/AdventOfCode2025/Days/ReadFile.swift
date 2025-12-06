//
//  ReadFile.swift
//  AdventOfCode
//
//  Created by Matt Slocum on 12/1/24.
//
import Foundation


func readTextFile(file: String) throws -> String {
    // Try multiple bundle locations for compatibility with app execution and previews
    let bundlesToTry: [Bundle] = [.main, Bundle(for: BundleFinder.self)]
    
    for bundle in bundlesToTry {
        if let fileURL = bundle.url(forResource: file, withExtension: "txt") {
            do {
                return try String(contentsOf: fileURL, encoding: .utf8).trimmingCharacters(in: .whitespacesAndNewlines)
            } catch {
                // Continue to next bundle if reading fails
                continue
            }
        }
    }
    
    throw NSError(domain: "File not found", code: 404, userInfo: [NSLocalizedDescriptionKey: "Could not find \(file).txt in bundle"])
}

// Helper class to find the bundle
private class BundleFinder {}

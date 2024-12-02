//
//  ReadFile.swift
//  AdventOfCode
//
//  Created by Matt Slocum on 12/1/24.
//
import Foundation


func readTextFile(file: String) throws -> String {
    guard let fileURL = Bundle.main.url(forResource: file, withExtension: "txt") else {
        throw NSError(domain: "File not found", code: 404, userInfo: nil)
    }

    return try String(contentsOf: fileURL, encoding: .utf8).trimmingCharacters(in: .whitespacesAndNewlines)
}

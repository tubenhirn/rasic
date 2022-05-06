const commitAnalyzerOptions = {
  preset: "angular",
  releaseRules: [
    { type: "docs", release: false },
    { type: "refactor", release: "patch" },
    { type: "chore", release: "patch" },
  ],
};
const releaseNotesGeneratorOptions = {
  parserOpts: {
    headerPattern: /^(\w*)(?:\((.*)\))?: (.*)$/,
    headerCorrespondence: ["type", "scope", "subject"],
    noteKeywords: ["BREAKING CHANGE"],
    revertPattern: /^(?:Revert|revert:)\s"?([\s\S]+?)"?\s*This reverts commit (\w*)\./i,
    revertCorrespondence: ["header", "hash"],
  },
  writerOpts: {
    transform: (commit, context) => {
      let discard = true;
      const issues = [];

      commit.notes.forEach((note) => {
        note.title = "BREAKING CHANGES";
        discard = false;
      });

      if (commit.type === "feat") {
        commit.type = "Features";
      } else if (commit.type === "fix") {
        commit.type = "Bug Fixes";
      } else if (commit.type === "perf") {
        commit.type = "Performance Improvements";
      } else if (commit.type === "revert" || commit.revert) {
        commit.type = "Reverts";
      } else if (commit.type === "docs") {
        commit.type = "Documentation";
      } else if (commit.type === "style") {
        commit.type = "Styles";
      } else if (commit.type === "refactor") {
        commit.type = "Code Refactoring";
      } else if (commit.type === "chore") {
        commit.type = "Miscellaneous Chores";
      } else if (commit.type === "test") {
        commit.type = "Tests";
      } else if (commit.type === "build") {
        commit.type = "Build System";
      } else if (commit.type === "ci") {
        commit.type = "Continuous Integration";
      } else {
        return;
      }

      if (commit.scope === "*") {
        commit.scope = "";
      }

      if (typeof commit.hash === "string") {
        commit.shortHash = commit.hash.substring(0, 7);
      }

      // remove references that already appear in the subject
      commit.references = commit.references.filter((reference) => {
        if (issues.indexOf(reference.issue) === -1) {
          return true;
        }

        return false;
      });

      // DO NOT REMOVE
      return commit;
    },
  },
};

module.exports = {
  extends: "/opt/base.config.js",
  branches: ["main"],
  plugins: [
    ["@semantic-release/commit-analyzer", commitAnalyzerOptions],
    ["@semantic-release/release-notes-generator", releaseNotesGeneratorOptions],
    [
      "@semantic-release/changelog",
      {
        changelogFile: "CHANGELOG.md",
      },
    ],
    [
      "@google/semantic-release-replace-plugin",
      {
        replacements: [
          {
            files: ["version"],
            from: "v.*",
            to: "v${nextRelease.version}",
            results: [
              {
                file: "version",
                hasChanged: true,
                numMatches: 1,
                numReplacements: 1,
              },
            ],
            countMatches: true,
          },
        ],
      },
    ],
    [
      "@semantic-release/git",
      {
        assets: ["version", "CHANGELOG.md"],
      },
    ],
  ],
};

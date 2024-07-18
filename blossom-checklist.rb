class BlossomChecklist < Formula
  desc "A pre-push checklist tool"
  homepage "https://github.com/avomakesart/homebrew-blossom-checklist"
  url "https://github.com/avomakesart/homebrew-blossom-checklist/archive/refs/tags/v1.0.1.tar.gz"
  sha256 "6242e2195cb5a834876f845fb5aa3a57257f464c4b9a2e83fb70d4aa31e5f992"
  version "1.0.1"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "blossom-checklist"
    bin.install "blossom-checklist"
  end

  test do
    system "#{bin}/blossom-checklist", "--version"
  end
end

class Checklist < Formula
  desc "A pre-push checklist tool"
  homepage "https://github.com/avomakesart/homebrew-blossom-checklist"
  url "https://github.com/avomakesart/homebrew-blossom-checklist/archive/refs/tags/blossom-first.tar.gz"
  sha256 "336063f228b5bda3fd315350a7b47466527ded221c315c973d6012057a324b4d"
  version "1.0.1"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "checklist"
    bin.install "blossom-checklist"
  end

  test do
    system "#{bin}/blossom-checklist", "--version"
  end
end

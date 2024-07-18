class BlossomChecklist < Formula
  desc "A pre-push checklist tool"
  homepage "https://github.com/avomakesart/homebrew-blossom-checklist"
  url "https://github.com/avomakesart/homebrew-blossom-checklist/archive/refs/tags/v1.1.0.tar.gz"
  sha256 "726b312c46678495a14876848723887eec9de3a0a74e184f53f705b255bbe836"
  version "1.1.0"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "blossom-checklist"
    bin.install "blossom-checklist"
  end

  test do
    system "#{bin}/blossom-checklist", "--version"
  end
end

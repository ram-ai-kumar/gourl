class Gourl < Formula
  desc "Smart CLI utility for managing and quickly opening project URLs"
  homepage "https://github.com/ram-ai-kumar/gourl"
  url "https://github.com/ram-ai-kumar/gourl.git",
    tag: "v1.0.0",
    revision: "HEAD"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "-o", bin/"gourl"
  end

  test do
    output = shell_output("#{bin}/gourl help")
    assert_match("gourl - Project URL Manager", output)
    assert_match("Usage:", output)
  end
end

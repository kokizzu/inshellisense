class Inshellisense < Formula
  desc "IDE style command line auto complete"
  homepage "https://github.com/microsoft/inshellisense"
  url "https://registry.npmjs.org/@microsoft/inshellisense/-/inshellisense-0.0.1.tgz"
  sha256 "2d1ce2ddfc7340f07990b34b14bc913d3ccf3f50c43e0a9ebee932ce05919f9b"
  license "MIT"

  depends_on "node"

  def install
    system "npm", "install", *std_npm_args
    bin.install_symlink libexec.glob("bin/*")
  end

  test do
    system bin/"inshellisense", "--version"
  end
end
	
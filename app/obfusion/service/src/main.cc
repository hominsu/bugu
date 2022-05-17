// MIT License
//
// Copyright (c) 2022. HominSu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

//
// Created by HominSu on 2022/5/16.
//

#include "obfusion/include/obfusion.h"
#include <cstdio>
#ifdef _WIN32
#include <windows.h>
#endif

#include "api/obfusion/service/v1/cpp/bugu_obfusion.pb.h"

bool read_file(const char *filename, u8 **data, u32 *size) {
  bool ret = false;
  FILE *f = fopen(filename, "rb");
  if (f) {
    fseek(f, 0, SEEK_END);
    u32 fsize = ftell(f);
    fseek(f, 0, SEEK_SET);

    u8 *bindata = new u8[fsize];
    if (fread(bindata, 1, fsize, f) == fsize) {
      *data = bindata;
      *size = fsize;
      ret = true;
    } else
      delete[] bindata;

    fclose(f);
  }
  return ret;
}

bool write_file(const char *fname, u8 *data, u32 size) {
  bool ret = false;
  FILE *f = fopen(fname, "wb");
  if (f) {
    if (fwrite(data, 1, size, f) == size)
      ret = true;

    fclose(f);
  }
  return ret;
}

int main() {
  obfusion obf;

  u8 *shell_data = nullptr;
  u32 shell_size;

  if (read_file("../res/exec_calc.bin", &shell_data, &shell_size)) {
    // shellcode binary data is present from offset 184 to 189 ("calc\0" string)
    // we need to supply this information so the obfuscator doesn't try to disassemble binary data
    std::vector<std::pair<u32, u32>> data_ranges;
    data_ranges.push_back(std::make_pair<u32, u32>(184, 189));

    // initial setup of obfuscation parameters
    obf.set_obf_steps(5,
                      20); // we want the obfuscated instructions to be divided into 5-20 separate instructions (approximately) (def. 3-8)
    obf.set_jmp_perc(20); // there should be 20% chance of inserting a jump at every instruction of the obfuscated output (def. 5%)

    obf.load32(shell_data, shell_size, data_ranges);

    delete[] shell_data;
    shell_data = nullptr;

    printf("obfuscating...\n");

    // obfuscate with random seed of 0xCAFEBABE and 3 obfuscation passes
    obf.obfuscate(0xCAFEBABE, 3);

    // dump the obfuscated shellcode into the buffer
    obf.dump(&shell_data, &shell_size);

    // save obfuscated shellcode to binary file
    if (write_file("../res/output.bin", shell_data, shell_size))
      printf("saved.\n");

#ifdef _WIN32
    DWORD prot;
        VirtualProtect(shell_data, shell_size, PAGE_EXECUTE_READWRITE, &prot);

        // execute obfuscated shellcode to test if runs properly (should execute calc.exe)
        ((void(*)())shell_data)();
#endif

    delete[] shell_data;
  }

  return 0;
}

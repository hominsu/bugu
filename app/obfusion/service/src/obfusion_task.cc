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
// Created by Homin Su on 2022/5/19.
//

#include "obfusion_task.h"
#include "obfusion/include/obfusion.h"
#include "data.h"

void bugu::ObfusionTask::Main() {
  obfusion obf;
  // shellcode binary data is present from offset 184 to 189 ("calc\0" string)
  // we need to supply this information so the obfuscator doesn't try to disassemble binary data
  ::std::vector<::std::pair<u32, u32>> data_ranges;
  data_ranges.push_back(::std::make_pair<u32, u32>(184, 189));

  // initial setup of obfuscation parameters
  obf.set_obf_steps(5,
                    20); // we want the obfuscated instructions to be divided into 5-20 separate instructions (approximately) (def. 3-8)
  obf.set_jmp_perc(20); // there should be 20% chance of inserting a jump at every instruction of the obfuscated output (def. 5%)

  obf.load32(data_->data(), data_->size(), data_ranges);

  // obfuscate with random seed of 0xCAFEBABE and 3 obfuscation passes
  obf.obfuscate(0xCAFEBABE, 3);

  u8 *shell_data = nullptr;
  u32 shell_size;
  // dump the obfuscated shellcode into the buffer
  obf.dump(&shell_data, &shell_size);

  auto data = ::bugu::Data::Make(memory_resource_);
  auto buffer = static_cast<char *>(data->New(shell_size));

  memcpy(buffer, shell_data, shell_size);
  delete[] shell_data;

  set_return(data);
}

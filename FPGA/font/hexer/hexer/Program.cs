using Coshx.IntelHexParser;
using System.Drawing;
using System;
using System.IO;
using System.Linq;
using System.Text;

namespace hexer
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var serializer = new Serializer();
            var bitmap = (Bitmap)Bitmap.FromFile(@"S:\Repos\MCPC\FPGA\font\ExportedFont.bmp");

            var result = new byte[128*16*8];
            
            for (var char_y = 0; char_y < 8; char_y++)
            {
                for (var char_x = 0; char_x < 16; char_x++)
                {
                    var x_from = char_x * 8;
                    var y_from = char_y * 16;

                    var asciiCode = char_x + char_y * 16;
                    Console.WriteLine("Preparing character: " + asciiCode);

                    for (var col = 0; col < 8; col++)
                    {
                        for (var row = 0; row < 16; row++)
                        {
                            var pixel = bitmap.GetPixel(x_from + col, y_from + row);
                            result[(asciiCode << 7) | (col << 4) | row] = (byte)(pixel.R == 0 ? 0 : 0x1);
                        }
                    }
                }
            }

            Console.WriteLine("Packing...");
            var builder = new StringBuilder();
            
            for (int i = 0; i < result.Length; i++)
            {
                if (result[i] == 0)
                {
                    builder.AppendLine("0");
                }
                else
                {
                    builder.AppendLine("1");
                }
            }

            Console.WriteLine("Writing to file...");
            File.WriteAllText(@"S:\Repos\MCPC\FPGA\src\VGA\font_data.raw", builder.ToString());

            Console.WriteLine("Done!");
            Console.ReadKey(true);
        }
    }
}

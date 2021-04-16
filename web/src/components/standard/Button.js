import styled from 'styled-components';
import { Colors } from '../../utils/Colors';

export const InputButton = styled.input`
  padding: 6px 12px;
  cursor: pointer;
  background-color: ${Colors.GrayD};

  border-radius: 3px;

  :hover {
    background-color: ${Colors.GrayC};
  }
`;

export const Button = styled.button`
  padding: 6px 12px;
  cursor: pointer;
  background-color: ${Colors.GrayD};

  border-radius: 3px;

  :hover {
    background-color: ${Colors.GrayC};
  }
`;
